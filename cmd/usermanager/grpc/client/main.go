package main

import (
	"context"
	"fmt"
	"strings"

	usergrpcClient "usermanager/grpc/client"
	"usermanager/internal/apperrors"
	"usermanager/internal/config"
	"usermanager/internal/domain/model"
	"usermanager/internal/infrastructure/logger"
	"usermanager/internal/utils"

	"github.com/google/uuid"
)

const dotEnv = "./configs/.env"

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(dotEnv)
	if err != nil {
		logger.Fatal(err)
	}

	connectionStr := fmt.Sprintf("localhost:%s", cfg.PortGrpcClient)
	grpcService, err := usergrpcClient.NewGrpcService(connectionStr)
	if err != nil {
		logger.Fatal(err)
	}

	ctx := context.Background()

	user, err := grpcService.CreateUser(ctx, &model.User{
		Nickname:  "testNickname",
		LastName:  "testLastName",
		FirstName: "testFirstName",
		Email:     "testEmail@test.test",
	})
	if err != nil {
		logger.Fatal(err)
	}

	user, err = grpcService.UpdateUser(ctx, &model.User{
		UserID:    user.UserID,
		Nickname:  "testNickname2",
		LastName:  "testLastName2",
		FirstName: "testFirstName2",
		Email:     "testEmail2@test.test",
	})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(user)

	err = grpcService.DeleteUser(ctx, &user.UserID)
	if err != nil {
		logger.Fatal(err)
	}

	users, err := grpcService.GetUsers(ctx, &utils.PaginationQuery{
		Size: 10,
		Page: 1,
	})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Printf("%+v\n", users)

	users, err = grpcService.GetUsersByPaginationQuery(ctx, &utils.PaginationQuery{
		Size: 10,
		Page: 1,
	})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(users)

	userUUIDstr := "b760bddd-3903-4b6b-8bb5-826a82c8761e"
	userUUID, err := uuid.Parse(userUUIDstr)
	if err != nil {
		logger.Fatal(err)
	}
	userLoaded, err := grpcService.GetUser(ctx, userUUID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(userLoaded)

	userLoadedByID, err := grpcService.GetUserByID(ctx, userLoaded.UserID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(userLoadedByID)

	userUUIDstr = "2140e95d-2280-497a-a540-54bbc75636b6"
	userUUID, err = uuid.Parse(userUUIDstr)
	if err != nil {
		logger.Fatal(err)
	}
	userLoadedForVote, err := grpcService.GetUser(ctx, userUUID)
	if err != nil {
		logger.Fatal(err)
	}

	vote, userVote, err := grpcService.VoteUser(ctx, &model.Vote{
		Vote:          1,
		CreatedUserID: userLoadedByID.UserID,
	}, &model.UserVote{
		UserID: userLoadedForVote.UserID,
	})
	if err != nil {
		if !strings.Contains(err.Error(), apperrors.UserGrpcControllerVoteUser.Code) {
			logger.Fatal(err)
		}
	}
	fmt.Println(vote)
	fmt.Println(userVote)

	vote, userVote, err = grpcService.VoteUserWithdraw(ctx, &model.Vote{
		Vote:          1,
		CreatedUserID: userLoadedByID.UserID,
	}, &model.UserVote{
		UserID: userLoadedForVote.UserID,
	})

	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(vote)
	fmt.Println(userVote)

	vote, userVote, err = grpcService.FindExistVoting(ctx, &userLoadedByID.UserID, &userLoadedForVote.UserID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(vote)
	fmt.Println(userVote)

	votes, err := grpcService.FindVotesForUser(ctx, &userLoadedByID.UserID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(votes)

	usersVotes := &model.Users{
		Users: []*model.User{
			{
				UserID: userLoadedByID.UserID,
			},
		},
	}
	loadVotesToUsers, err := grpcService.LoadVotesToUsers(ctx, usersVotes)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(loadVotesToUsers)

	vote, err = grpcService.GetLastVoteForUser(ctx, &userLoadedByID.UserID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(vote)

	vote, userVote, err = grpcService.Vote(ctx, &model.Vote{
		Vote:          1,
		CreatedUserID: userLoadedByID.UserID,
	}, &model.UserVote{
		UserID: userLoadedForVote.UserID,
	})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(vote)
	fmt.Println(userVote)
}
