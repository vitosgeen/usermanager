package server

import (
	"context"

	grpcUsermanager "usermanager/grpc"
	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/usecase/usecase"
	"usermanager/internal/utils"

	"github.com/google/uuid"
)

type UserManagerGrpcController struct {
	userUscase usecase.IUserUsecase
	grpcUsermanager.UnimplementedUserUsecaseServer
}

func NewUserManagerGrpcController(userUscase usecase.IUserUsecase) *UserManagerGrpcController {
	return &UserManagerGrpcController{userUscase: userUscase}
}

func (umg *UserManagerGrpcController) CreateUser(ctx context.Context, userRequest *grpcUsermanager.CreateUserRequest) (*grpcUsermanager.CreateUserResponse, error) {
	user := &model.User{
		Nickname:  userRequest.User.Nickname,
		FirstName: userRequest.User.FirstName,
		LastName:  userRequest.User.LastName,
		Email:     userRequest.User.Email,
		Password:  userRequest.User.Password,
		IsPublic:  userRequest.User.IsPublic,
		Role:      userRequest.User.UserRole,
	}
	createdUser, err := umg.userUscase.CreateUser(ctx, user)
	if err != nil {
		return nil, apperrors.UserGrpcControllerCreateUserError.AppendMessage(err)
	}

	return &grpcUsermanager.CreateUserResponse{
		User: marshalUser(createdUser),
	}, nil
}

func (umg *UserManagerGrpcController) UpdateUser(ctx context.Context, userRequest *grpcUsermanager.UpdateUserRequest) (*grpcUsermanager.UpdateUserResponse, error) {
	userId, err := uuid.Parse(userRequest.User.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerUpdateUserUuidParse.AppendMessage(err)
	}
	user := &model.User{
		UserID:    userId,
		Nickname:  userRequest.User.Nickname,
		FirstName: userRequest.User.FirstName,
		LastName:  userRequest.User.LastName,
		Email:     userRequest.User.Email,
		Password:  userRequest.User.Password,
		IsPublic:  userRequest.User.IsPublic,
		Role:      userRequest.User.UserRole,
	}
	updatedUser, err := umg.userUscase.UpdateUser(ctx, user)
	if err != nil {
		return nil, apperrors.UserGrpcControllerUpdateUserUpdateUser.AppendMessage(err)
	}

	return &grpcUsermanager.UpdateUserResponse{
		User: marshalUser(updatedUser),
	}, nil
}

func (umg *UserManagerGrpcController) DeleteUser(ctx context.Context, userRequest *grpcUsermanager.DeleteUserRequest) (*grpcUsermanager.DeleteUserResponse, error) {
	userId, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerDeleteUserUuidParse.AppendMessage(err)
	}
	err = umg.userUscase.DeleteUser(ctx, &userId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerDeleteUser.AppendMessage(err)
	}

	return &grpcUsermanager.DeleteUserResponse{}, nil
}

func (umg *UserManagerGrpcController) GetUsers(ctx context.Context, usersRequest *grpcUsermanager.GetUsersRequest) (*grpcUsermanager.GetUsersResponse, error) {
	paginationQuery := &utils.PaginationQuery{
		Size:    int(usersRequest.PaginationQuery.Size),
		Page:    int(usersRequest.PaginationQuery.Page),
		OrderBy: usersRequest.PaginationQuery.OrderBy,
	}
	users, err := umg.userUscase.GetUsers(ctx, paginationQuery)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUsers.AppendMessage(err)
	}

	grpcUsers := &grpcUsermanager.Users{
		Page:    int32(users.Page),
		Users:   []*grpcUsermanager.User{},
		HasMore: users.HasMore,
	}
	for _, user := range users.Users {
		grpcUsers.Users = append(grpcUsers.Users, marshalUser(user))
	}

	return &grpcUsermanager.GetUsersResponse{
		Users: grpcUsers,
	}, nil
}

func (umg *UserManagerGrpcController) GetUsersByPaginationQuery(ctx context.Context, usersRequest *grpcUsermanager.GetUsersByPaginationQueryRequest) (*grpcUsermanager.GetUsersByPaginationQueryResponse, error) {
	paginationQuery := &utils.PaginationQuery{
		Size:    int(usersRequest.PaginationQuery.Size),
		Page:    int(usersRequest.PaginationQuery.Page),
		OrderBy: usersRequest.PaginationQuery.OrderBy,
	}
	users, err := umg.userUscase.GetUsersByPaginationQuery(ctx, paginationQuery)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUsersByPaginationQuery.AppendMessage(err)
	}

	grpcUsers := &grpcUsermanager.Users{
		Page:    int32(users.Page),
		Users:   []*grpcUsermanager.User{},
		HasMore: users.HasMore,
	}
	for _, user := range users.Users {
		grpcUsers.Users = append(grpcUsers.Users, marshalUser(user))
	}

	return &grpcUsermanager.GetUsersByPaginationQueryResponse{
		Users: grpcUsers,
	}, nil
}

func (umg *UserManagerGrpcController) GetUser(ctx context.Context, userRequest *grpcUsermanager.GetUserRequest) (*grpcUsermanager.GetUserResponse, error) {
	userId, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUserUuidParse.AppendMessage(err)
	}
	user, err := umg.userUscase.GetUser(ctx, userId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUser.AppendMessage(err)
	}

	return &grpcUsermanager.GetUserResponse{
		User: marshalUser(user),
	}, nil
}

func (umg *UserManagerGrpcController) GetUserByID(ctx context.Context, userRequest *grpcUsermanager.GetUserByIDRequest) (*grpcUsermanager.GetUserByIDResponse, error) {
	userId, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUserByIDUuidParse.AppendMessage(err)
	}
	user, err := umg.userUscase.GetUserByID(ctx, userId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUserByID.AppendMessage(err)
	}

	return &grpcUsermanager.GetUserByIDResponse{
		User: marshalUser(user),
	}, nil
}

func (umg *UserManagerGrpcController) GetUserByNickname(ctx context.Context, userRequest *grpcUsermanager.GetUserByNicknameRequest) (*grpcUsermanager.GetUserByNicknameResponse, error) {
	user, err := umg.userUscase.GetUserByNickname(ctx, userRequest.Nickname)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetUserByNickname.AppendMessage(err)
	}

	return &grpcUsermanager.GetUserByNicknameResponse{
		User: marshalUser(user),
	}, nil
}

func (umg *UserManagerGrpcController) CheckUserByNickname(ctx context.Context, userRequest *grpcUsermanager.CheckUserByNicknameRequest) (*grpcUsermanager.CheckUserByNicknameResponse, error) {
	isExist, err := umg.userUscase.CheckUserByNickname(ctx, marshalGrpcUserToUser(userRequest.User))
	if err != nil {
		return nil, apperrors.UserGrpcControllerCheckUserByNickname.AppendMessage(err)
	}

	return &grpcUsermanager.CheckUserByNicknameResponse{
		IsExist: isExist,
	}, nil
}

func (umg *UserManagerGrpcController) VoteUser(ctx context.Context, userRequest *grpcUsermanager.VoteUserRequest) (*grpcUsermanager.VoteUserResponse, error) {
	vote := marshalGrpcVoteToVote(userRequest.Vote)
	userVote := marshalGrpcUserVoteToUserVote(userRequest.UserVote)
	voteUser, userVote, err := umg.userUscase.VoteUser(ctx, vote, userVote)
	if err != nil {
		return nil, apperrors.UserGrpcControllerVoteUser.AppendMessage(err)
	}

	return &grpcUsermanager.VoteUserResponse{
		Vote:     marshalVoteToGrpcVote(voteUser),
		UserVote: marshalUserVoteToGrpcUserVote(userVote),
	}, nil
}

func (umg *UserManagerGrpcController) VoteUserWithdraw(ctx context.Context, userRequest *grpcUsermanager.VoteUserWithdrawRequest) (*grpcUsermanager.VoteUserWithdrawResponse, error) {
	vote := marshalGrpcVoteToVote(userRequest.Vote)
	userVote := marshalGrpcUserVoteToUserVote(userRequest.UserVote)
	voteUser, userVote, err := umg.userUscase.VoteUserWithdraw(ctx, vote, userVote)
	if err != nil {
		return nil, apperrors.UserGrpcControllerVoteUser.AppendMessage(err)
	}

	return &grpcUsermanager.VoteUserWithdrawResponse{
		Vote:     marshalVoteToGrpcVote(voteUser),
		UserVote: marshalUserVoteToGrpcUserVote(userVote),
	}, nil
}

func (umg *UserManagerGrpcController) FindExistVoting(ctx context.Context, userRequest *grpcUsermanager.FindExistVotingRequest) (*grpcUsermanager.FindExistVotingResponse, error) {
	userID, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerFindExistVoting.AppendMessage(err)
	}
	voterUserID, err := uuid.Parse(userRequest.VoterUserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerFindExistVoting.AppendMessage(err)
	}

	voteUser, userVote, err := umg.userUscase.FindExistVoting(ctx, &userID, &voterUserID)
	if err != nil {
		return nil, apperrors.UserGrpcControllerFindExistVoting.AppendMessage(err)
	}

	if voteUser == nil && userVote == nil {
		return &grpcUsermanager.FindExistVotingResponse{}, nil
	}

	return &grpcUsermanager.FindExistVotingResponse{
		Vote:     marshalVoteToGrpcVote(voteUser),
		UserVote: marshalUserVoteToGrpcUserVote(userVote),
	}, nil
}

func (umg *UserManagerGrpcController) FindVotesForUser(ctx context.Context, userRequest *grpcUsermanager.FindVotesForUserRequest) (*grpcUsermanager.FindVotesForUserResponse, error) {
	userID, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerFindVotesForUserUuidParse.AppendMessage(err)
	}

	votes, err := umg.userUscase.FindVotesForUser(ctx, &userID)
	if err != nil {
		return nil, apperrors.UserGrpcControllerFindVotesForUserFindVotesForUser.AppendMessage(err)
	}

	return &grpcUsermanager.FindVotesForUserResponse{
		Votes: marshalVotes(votes),
	}, nil
}

func (umg *UserManagerGrpcController) LoadVotesToUsers(ctx context.Context, userRequest *grpcUsermanager.LoadVotesToUsersRequest) (*grpcUsermanager.LoadVotesToUsersResponse, error) {
	usersMarshaled := &model.Users{}
	for _, user := range userRequest.Users.Users {
		marshalUser := marshalGrpcUserToUser(user)
		usersMarshaled.Users = append(usersMarshaled.Users, marshalUser)
	}

	users, err := umg.userUscase.LoadVotesToUsers(ctx, usersMarshaled)
	if err != nil {
		return nil, apperrors.UserGrpcControllerLoadVotesToUsers.AppendMessage(err)
	}

	return &grpcUsermanager.LoadVotesToUsersResponse{
		Users: marshalGrpcUsersToUsers(users),
	}, nil
}

func (umg *UserManagerGrpcController) GetLastVoteForUser(ctx context.Context, userRequest *grpcUsermanager.GetLastVoteForUserRequest) (*grpcUsermanager.GetLastVoteForUserResponse, error) {
	userID, err := uuid.Parse(userRequest.UserId)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetLastVoteForUserUuidParse.AppendMessage(err)
	}

	vote, err := umg.userUscase.GetLastVoteForUser(ctx, &userID)
	if err != nil {
		return nil, apperrors.UserGrpcControllerGetLastVoteForUser.AppendMessage(err)
	}

	return &grpcUsermanager.GetLastVoteForUserResponse{
		Vote: marshalVoteToGrpcVote(vote),
	}, nil
}

func (umg *UserManagerGrpcController) Vote(ctx context.Context, userRequest *grpcUsermanager.VoteRequest) (*grpcUsermanager.VoteResponse, error) {
	vote := marshalGrpcVoteToVote(userRequest.Vote)
	userVote := marshalGrpcUserVoteToUserVote(userRequest.UserVote)
	voteUser, userVoteUser, err := umg.userUscase.Vote(ctx, vote, userVote)
	if err != nil {
		return nil, apperrors.UserGrpcControllerVote.AppendMessage(err)
	}

	return &grpcUsermanager.VoteResponse{
		Vote:     marshalVoteToGrpcVote(voteUser),
		UserVote: marshalUserVoteToGrpcUserVote(userVoteUser),
	}, nil
}

func marshalGrpcUsersToUsers(users []*model.User) []*grpcUsermanager.User {
	var grpcUsers []*grpcUsermanager.User
	for _, user := range users {
		grpcUsers = append(grpcUsers, marshalUser(user))
	}
	return grpcUsers
}

func marshalVoteToGrpcVote(vote *model.Vote) *grpcUsermanager.Vote {
	return &grpcUsermanager.Vote{
		VoteId:        vote.VoteID,
		Vote:          int32(vote.Vote),
		CreatedUserId: vote.CreatedUserID.String(),
		CreatedAt:     vote.CreatedAt.String(),
	}
}

func marshalGrpcVoteToVote(vote *grpcUsermanager.Vote) *model.Vote {
	timeCreated, _ := utils.ParseStrToTime(vote.CreatedAt)
	return &model.Vote{
		VoteID:        vote.VoteId,
		Vote:          int(vote.Vote),
		CreatedUserID: uuid.MustParse(vote.CreatedUserId),
		CreatedAt:     &timeCreated,
	}
}

func marshalGrpcUserVoteToUserVote(userVote *grpcUsermanager.UserVote) *model.UserVote {
	return &model.UserVote{
		ID:     userVote.Id,
		UserID: uuid.MustParse(userVote.UserId),
		VoteID: userVote.VoteId,
	}
}

func marshalUserVoteToGrpcUserVote(userVote *model.UserVote) *grpcUsermanager.UserVote {
	return &grpcUsermanager.UserVote{
		Id:     userVote.ID,
		UserId: userVote.UserID.String(),
		VoteId: userVote.VoteID,
	}
}

func marshalGrpcUserToUser(user *grpcUsermanager.User) *model.User {
	return &model.User{
		UserID:    uuid.MustParse(user.UserId),
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsPublic:  user.IsPublic,
		Role:      user.UserRole,
		Votes:     marshalGrpcVotesToVotes(user.Votes),
	}
}

func marshalGrpcVotesToVotes(votes []*grpcUsermanager.Vote) []*model.Vote {
	var modelVotes []*model.Vote
	for _, vote := range votes {
		timeCreated, _ := utils.ParseStrToTime(vote.CreatedAt)
		modelVotes = append(modelVotes, &model.Vote{
			VoteID:        vote.VoteId,
			Vote:          int(vote.Vote),
			CreatedUserID: uuid.MustParse(vote.CreatedUserId),
			CreatedAt:     &timeCreated,
		})
	}
	return modelVotes
}

func marshalUser(user *model.User) *grpcUsermanager.User {
	return &grpcUsermanager.User{
		UserId:    user.UserID.String(),
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsPublic:  user.IsPublic,
		UserRole:  user.Role,
		Votes:     marshalVotes(user.Votes),
	}
}

func marshalVotes(votes []*model.Vote) []*grpcUsermanager.Vote {
	var grpcVotes []*grpcUsermanager.Vote
	for _, vote := range votes {
		grpcVotes = append(grpcVotes, &grpcUsermanager.Vote{
			VoteId:        vote.VoteID,
			Vote:          int32(vote.Vote),
			CreatedUserId: vote.CreatedUserID.String(),
			CreatedAt:     vote.CreatedAt.String(),
		})
	}
	return grpcVotes
}
