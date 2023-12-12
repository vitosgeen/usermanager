package server

import (
	"context"

	"usermanager/internal/domain/model"
	"usermanager/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (uum *UserUsecaseMock) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	args := uum.Called(ctx, user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	args := uum.Called(ctx, user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) DeleteUser(ctx context.Context, userID *uuid.UUID) error {
	args := uum.Called(ctx, userID)
	return args.Error(0)
}

func (uum *UserUsecaseMock) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	args := uum.Called(ctx, paginationQuery)
	return args.Get(0).(*model.Users), args.Error(1)
}

func (uum *UserUsecaseMock) GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	args := uum.Called(ctx, paginationQuery)
	return args.Get(0).(*model.Users), args.Error(1)
}

func (uum *UserUsecaseMock) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	args := uum.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	args := uum.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	args := uum.Called(ctx, nickname)
	return args.Get(0).(*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) CheckUserByNickname(ctx context.Context, user *model.User) (bool, error) {
	args := uum.Called(ctx, user)
	return args.Bool(0), args.Error(1)
}

func (uum *UserUsecaseMock) VoteUser(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	args := uum.Called(ctx, vote, userVote)
	return args.Get(0).(*model.Vote), args.Get(1).(*model.UserVote), args.Error(2)
}

func (uum *UserUsecaseMock) VoteUserWithdraw(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	args := uum.Called(ctx, vote, userVote)
	return args.Get(0).(*model.Vote), args.Get(1).(*model.UserVote), args.Error(2)
}

func (uum *UserUsecaseMock) FindExistVoting(ctx context.Context, userID *uuid.UUID, voterUserID *uuid.UUID) (*model.Vote, *model.UserVote, error) {
	args := uum.Called(ctx, userID, voterUserID)
	return args.Get(0).(*model.Vote), args.Get(1).(*model.UserVote), args.Error(2)
}

func (uum *UserUsecaseMock) FindVotesForUser(ctx context.Context, userID *uuid.UUID) ([]*model.Vote, error) {
	args := uum.Called(ctx, userID)
	return args.Get(0).([]*model.Vote), args.Error(1)
}

func (uum *UserUsecaseMock) LoadVotesToUsers(ctx context.Context, users *model.Users) ([]*model.User, error) {
	args := uum.Called(ctx, users)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (uum *UserUsecaseMock) GetLastVoteForUser(ctx context.Context, userID *uuid.UUID) (*model.Vote, error) {
	args := uum.Called(ctx, userID)
	return args.Get(0).(*model.Vote), args.Error(1)
}

func (uum *UserUsecaseMock) Vote(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	args := uum.Called(ctx, vote, userVote)
	return args.Get(0).(*model.Vote), args.Get(1).(*model.UserVote), args.Error(2)
}
