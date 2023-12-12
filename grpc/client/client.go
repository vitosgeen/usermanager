package client

import (
	"context"
	"fmt"

	grpcUsermanager "usermanager/grpc"
	"usermanager/internal/domain/model"
	"usermanager/internal/usecase/usecase"
	"usermanager/internal/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type grpcService struct {
	client grpcUsermanager.UserUsecaseClient
}

func NewGrpcService(connectionStr string) (usecase.IUserUsecase, error) {
	conn, err := grpc.Dial(connectionStr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{
		client: grpcUsermanager.NewUserUsecaseClient(conn),
	}, nil
}

func (s *grpcService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	createUserRequest := marshalUserToCreateUserRequest(user)
	createUserResponse, err := s.client.CreateUser(ctx, createUserRequest)
	if err != nil {
		return nil, err
	}
	user, err = marshalUserReponseToUser(createUserResponse)
	return user, nil
}

func (s *grpcService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	userUpdateRequest := marshalUserToUpdateUserRequest(user)
	userUpdateResponse, err := s.client.UpdateUser(ctx, userUpdateRequest)
	if err != nil {
		return nil, err
	}
	user, err = marshalUpdateUserResponseToUser(userUpdateResponse)
	return user, nil
}

func (s *grpcService) DeleteUser(ctx context.Context, userID *uuid.UUID) error {
	deleteUserRequest := &grpcUsermanager.DeleteUserRequest{
		UserId: userID.String(),
	}

	_, err := s.client.DeleteUser(ctx, deleteUserRequest)
	return err
}

func (s *grpcService) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	getUsersRequest := &grpcUsermanager.GetUsersRequest{
		PaginationQuery: &grpcUsermanager.PaginationQuery{
			Page: int32(paginationQuery.Page),
			Size: int32(paginationQuery.Size),
		},
	}

	getUsersResponse, err := s.client.GetUsers(ctx, getUsersRequest)
	if err != nil {
		return nil, err
	}

	users := &model.Users{}
	for _, user := range getUsersResponse.Users.Users {
		uuid, err := uuid.Parse(user.UserId)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &model.User{
			UserID:    uuid,
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		})
	}

	return users, nil
}

func (s *grpcService) GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	getUsersByPaginationQueryRequest := &grpcUsermanager.GetUsersByPaginationQueryRequest{
		PaginationQuery: &grpcUsermanager.PaginationQuery{
			Page: int32(paginationQuery.Page),
			Size: int32(paginationQuery.Size),
		},
	}

	getUsersByPaginationQueryResponse, err := s.client.GetUsersByPaginationQuery(ctx, getUsersByPaginationQueryRequest)
	if err != nil {
		return nil, err
	}

	users := &model.Users{}
	for _, user := range getUsersByPaginationQueryResponse.Users.Users {
		uuid, err := uuid.Parse(user.UserId)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &model.User{
			UserID:    uuid,
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		})
	}

	return users, nil
}

func (s *grpcService) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	getUserRequest := &grpcUsermanager.GetUserRequest{
		UserId: userID.String(),
	}

	getUserResponse, err := s.client.GetUser(ctx, getUserRequest)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(getUserResponse.User.UserId)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		UserID:    uuid,
		Nickname:  getUserResponse.User.Nickname,
		LastName:  getUserResponse.User.LastName,
		FirstName: getUserResponse.User.FirstName,
		Email:     getUserResponse.User.Email,
	}

	return user, nil
}

func (s *grpcService) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	getUserByIDRequest := &grpcUsermanager.GetUserByIDRequest{
		UserId: userID.String(),
	}

	getUserByIDResponse, err := s.client.GetUserByID(ctx, getUserByIDRequest)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(getUserByIDResponse.User.UserId)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		UserID:    uuid,
		Nickname:  getUserByIDResponse.User.Nickname,
		LastName:  getUserByIDResponse.User.LastName,
		FirstName: getUserByIDResponse.User.FirstName,
		Email:     getUserByIDResponse.User.Email,
	}

	return user, nil
}

func (s *grpcService) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	getUserByNicknameRequest := &grpcUsermanager.GetUserByNicknameRequest{
		Nickname: nickname,
	}

	getUserByNicknameResponse, err := s.client.GetUserByNickname(ctx, getUserByNicknameRequest)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(getUserByNicknameResponse.User.UserId)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		UserID:    uuid,
		Nickname:  getUserByNicknameResponse.User.Nickname,
		LastName:  getUserByNicknameResponse.User.LastName,
		FirstName: getUserByNicknameResponse.User.FirstName,
		Email:     getUserByNicknameResponse.User.Email,
	}

	return user, nil
}

func (s *grpcService) CheckUserByNickname(ctx context.Context, user *model.User) (bool, error) {
	checkUserByNicknameRequest := &grpcUsermanager.CheckUserByNicknameRequest{
		User: &grpcUsermanager.User{
			Nickname: user.Nickname,
		},
	}

	checkUserByNicknameResponse, err := s.client.CheckUserByNickname(ctx, checkUserByNicknameRequest)
	fmt.Println(checkUserByNicknameResponse)
	if err != nil {
		return false, err
	}

	return checkUserByNicknameResponse.IsExist, nil
}

func (s *grpcService) VoteUser(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	voteUserRequest := &grpcUsermanager.VoteUserRequest{
		Vote: &grpcUsermanager.Vote{
			VoteId:        vote.VoteID,
			Vote:          int32(vote.Vote),
			CreatedUserId: vote.CreatedUserID.String(),
		},
		UserVote: &grpcUsermanager.UserVote{
			Id:     userVote.VoteID,
			UserId: userVote.UserID.String(),
			VoteId: userVote.VoteID,
		},
	}

	voteUserResponse, err := s.client.VoteUser(ctx, voteUserRequest)
	if err != nil {
		return nil, nil, err
	}

	createdUserID, err := uuid.Parse(voteUserResponse.Vote.CreatedUserId)
	if err != nil {
		return nil, nil, err
	}
	vote = &model.Vote{
		VoteID:        voteUserResponse.Vote.VoteId,
		Vote:          int(voteUserResponse.Vote.Vote),
		CreatedUserID: createdUserID,
	}

	userID, err := uuid.Parse(voteUserResponse.UserVote.UserId)
	if err != nil {
		return nil, nil, err
	}
	userVote = &model.UserVote{
		ID:     voteUserResponse.UserVote.Id,
		UserID: userID,
		VoteID: voteUserResponse.UserVote.VoteId,
	}

	return vote, userVote, nil
}

func (s *grpcService) VoteUserWithdraw(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	voteUserWithdrawRequest := &grpcUsermanager.VoteUserWithdrawRequest{
		Vote: &grpcUsermanager.Vote{
			VoteId:        vote.VoteID,
			Vote:          int32(vote.Vote),
			CreatedUserId: vote.CreatedUserID.String(),
		},
		UserVote: &grpcUsermanager.UserVote{
			Id:     userVote.VoteID,
			UserId: userVote.UserID.String(),
			VoteId: userVote.VoteID,
		},
	}

	voteUserWithdrawResponse, err := s.client.VoteUserWithdraw(ctx, voteUserWithdrawRequest)
	if err != nil {
		return nil, nil, err
	}

	createdUserID, err := uuid.Parse(voteUserWithdrawResponse.Vote.CreatedUserId)
	if err != nil {
		return nil, nil, err
	}
	vote = &model.Vote{
		VoteID:        voteUserWithdrawResponse.Vote.VoteId,
		Vote:          int(voteUserWithdrawResponse.Vote.Vote),
		CreatedUserID: createdUserID,
	}

	userID, err := uuid.Parse(voteUserWithdrawResponse.UserVote.UserId)
	if err != nil {
		return nil, nil, err
	}
	userVote = &model.UserVote{
		ID:     voteUserWithdrawResponse.UserVote.Id,
		UserID: userID,
		VoteID: voteUserWithdrawResponse.UserVote.VoteId,
	}

	return vote, userVote, nil
}

func (s *grpcService) FindExistVoting(ctx context.Context, userID *uuid.UUID, voterUserID *uuid.UUID) (*model.Vote, *model.UserVote, error) {
	findExistVotingRequest := &grpcUsermanager.FindExistVotingRequest{
		UserId:      userID.String(),
		VoterUserId: voterUserID.String(),
	}

	findExistVotingResponse, err := s.client.FindExistVoting(ctx, findExistVotingRequest)
	if err != nil {
		return nil, nil, err
	}

	if findExistVotingResponse.Vote == nil && findExistVotingResponse.UserVote == nil {
		return nil, nil, nil
	}

	createdUserID, err := uuid.Parse(findExistVotingResponse.Vote.CreatedUserId)
	if err != nil {
		return nil, nil, err
	}
	vote := &model.Vote{
		VoteID:        findExistVotingResponse.Vote.VoteId,
		Vote:          int(findExistVotingResponse.Vote.Vote),
		CreatedUserID: createdUserID,
	}

	userVoteUserID, err := uuid.Parse(findExistVotingResponse.UserVote.UserId)
	if err != nil {
		return nil, nil, err
	}
	userVote := &model.UserVote{
		ID:     findExistVotingResponse.UserVote.Id,
		UserID: userVoteUserID,
		VoteID: findExistVotingResponse.UserVote.VoteId,
	}

	return vote, userVote, nil
}

func (s *grpcService) FindVotesForUser(ctx context.Context, userID *uuid.UUID) ([]*model.Vote, error) {
	findVotesForUserRequest := &grpcUsermanager.FindVotesForUserRequest{
		UserId: userID.String(),
	}

	findVotesForUserResponse, err := s.client.FindVotesForUser(ctx, findVotesForUserRequest)
	if err != nil {
		return nil, err
	}

	var votes []*model.Vote
	for _, vote := range findVotesForUserResponse.Votes {
		createdUserID, err := uuid.Parse(vote.CreatedUserId)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &model.Vote{
			VoteID:        vote.VoteId,
			Vote:          int(vote.Vote),
			CreatedUserID: createdUserID,
		})
	}

	return votes, nil
}

func (s *grpcService) LoadVotesToUsers(ctx context.Context, users *model.Users) ([]*model.User, error) {
	var usersToLoad []*grpcUsermanager.User
	for _, user := range users.Users {
		usersToLoad = append(usersToLoad, &grpcUsermanager.User{
			UserId:    user.UserID.String(),
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		})
	}

	loadVotesToUsersRequest := &grpcUsermanager.LoadVotesToUsersRequest{
		Users: &grpcUsermanager.Users{
			Users: usersToLoad,
		},
	}

	loadVotesToUsersResponse, err := s.client.LoadVotesToUsers(ctx, loadVotesToUsersRequest)
	if err != nil {
		return nil, err
	}

	var usersWithVotes []*model.User
	for _, user := range loadVotesToUsersResponse.Users {
		uuid, err := uuid.Parse(user.UserId)
		if err != nil {
			return nil, err
		}
		usersWithVotes = append(usersWithVotes, &model.User{
			UserID:    uuid,
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		})
	}

	return usersWithVotes, nil
}

func (s *grpcService) GetLastVoteForUser(ctx context.Context, userID *uuid.UUID) (*model.Vote, error) {
	getLastVoteForUserRequest := &grpcUsermanager.GetLastVoteForUserRequest{
		UserId: userID.String(),
	}

	getLastVoteForUserResponse, err := s.client.GetLastVoteForUser(ctx, getLastVoteForUserRequest)
	if err != nil {
		return nil, err
	}

	createdUserID, err := uuid.Parse(getLastVoteForUserResponse.Vote.CreatedUserId)
	if err != nil {
		return nil, err
	}
	vote := &model.Vote{
		VoteID:        getLastVoteForUserResponse.Vote.VoteId,
		Vote:          int(getLastVoteForUserResponse.Vote.Vote),
		CreatedUserID: createdUserID,
	}

	return vote, nil
}

func (s *grpcService) Vote(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	voteRequest := &grpcUsermanager.VoteRequest{
		Vote: &grpcUsermanager.Vote{
			VoteId:        vote.VoteID,
			Vote:          int32(vote.Vote),
			CreatedUserId: vote.CreatedUserID.String(),
		},
		UserVote: &grpcUsermanager.UserVote{
			Id:     userVote.VoteID,
			UserId: userVote.UserID.String(),
			VoteId: userVote.VoteID,
		},
	}

	voteResponse, err := s.client.Vote(ctx, voteRequest)
	if err != nil {
		return nil, nil, err
	}

	createdUserID, err := uuid.Parse(voteResponse.Vote.CreatedUserId)
	if err != nil {
		return nil, nil, err
	}
	vote = &model.Vote{
		VoteID:        voteResponse.Vote.VoteId,
		Vote:          int(voteResponse.Vote.Vote),
		CreatedUserID: createdUserID,
	}

	userID, err := uuid.Parse(voteResponse.UserVote.UserId)
	if err != nil {
		return nil, nil, err
	}
	userVote = &model.UserVote{
		ID:     voteResponse.UserVote.Id,
		UserID: userID,
		VoteID: voteResponse.UserVote.VoteId,
	}

	return vote, userVote, nil
}

func marshalUserToCreateUserRequest(user *model.User) *grpcUsermanager.CreateUserRequest {
	return &grpcUsermanager.CreateUserRequest{
		User: &grpcUsermanager.User{
			UserId:    user.UserID.String(),
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		},
	}
}

func marshalUserToUpdateUserRequest(user *model.User) *grpcUsermanager.UpdateUserRequest {
	return &grpcUsermanager.UpdateUserRequest{
		User: &grpcUsermanager.User{
			UserId:    user.UserID.String(),
			Nickname:  user.Nickname,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		},
	}
}

func marshalUpdateUserResponseToUser(userResponse *grpcUsermanager.UpdateUserResponse) (*model.User, error) {
	uuid, err := uuid.Parse(userResponse.User.UserId)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserID:    uuid,
		Nickname:  userResponse.User.Nickname,
		LastName:  userResponse.User.LastName,
		FirstName: userResponse.User.FirstName,
		Email:     userResponse.User.Email,
	}, nil
}

func marshalUserReponseToUser(userResponse *grpcUsermanager.CreateUserResponse) (*model.User, error) {
	uuid, err := uuid.Parse(userResponse.User.UserId)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserID:    uuid,
		Nickname:  userResponse.User.Nickname,
		LastName:  userResponse.User.LastName,
		FirstName: userResponse.User.FirstName,
		Email:     userResponse.User.Email,
	}, nil
}
