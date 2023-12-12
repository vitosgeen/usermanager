package usecase

import (
	"context"

	"usermanager/internal/domain/model"
	"usermanager/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (urm *UserRepositoryMock) FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	args := urm.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urm *UserRepositoryMock) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	args := urm.Called(ctx, nickname)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urm *UserRepositoryMock) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	args := urm.Called(ctx, paginationQuery)
	return args.Get(0).(*model.Users), args.Error(1)
}

func (urm *UserRepositoryMock) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	args := urm.Called(ctx, user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urm *UserRepositoryMock) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	args := urm.Called(ctx, user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urm *UserRepositoryMock) SoftDeleteUserByUserID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	args := urm.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urm *UserRepositoryMock) DeleteUserByUserID(ctx context.Context, userID *uuid.UUID) error {
	args := urm.Called(ctx, userID)
	return args.Error(1)
}

type UserRedisRepositoryMock struct {
	mock.Mock
}

func (urrm *UserRedisRepositoryMock) FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	args := urrm.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urrm *UserRedisRepositoryMock) SetFindUserByUUID(ctx context.Context, userID uuid.UUID, user *model.User) error {
	args := urrm.Called(ctx, userID, user)
	return args.Error(0)
}

func (urrm *UserRedisRepositoryMock) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	args := urrm.Called(ctx, nickname)
	return args.Get(0).(*model.User), args.Error(1)
}

func (urrm *UserRedisRepositoryMock) SetFindUserByNickname(ctx context.Context, nickname string, user *model.User) error {
	args := urrm.Called(ctx, nickname, user)
	return args.Error(0)
}

func (urrm *UserRedisRepositoryMock) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	args := urrm.Called(ctx, paginationQuery)
	return args.Get(0).(*model.Users), args.Error(1)
}

func (urrm *UserRedisRepositoryMock) SetGetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery, users *model.Users) error {
	args := urrm.Called(ctx, paginationQuery, users)
	return args.Error(0)
}

type VoteRedisRepositoryMock struct {
	mock.Mock
}

func (vrrm *VoteRedisRepositoryMock) FindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID) ([]*model.Vote, error) {
	args := vrrm.Called(ctx, voteUserID)
	return args.Get(0).([]*model.Vote), args.Error(1)
}

func (vrrm *VoteRedisRepositoryMock) SetFindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID, votes []*model.Vote) error {
	args := vrrm.Called(ctx, voteUserID, votes)
	return args.Error(0)
}

func (vrrm *VoteRedisRepositoryMock) FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error) {
	args := vrrm.Called(ctx, userIDs)
	return args.Get(0).([]*model.Vote), args.Error(1)
}

func (vrrm *VoteRedisRepositoryMock) SetFindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID, votes []*model.Vote) error {
	args := vrrm.Called(ctx, userIDs, votes)
	return args.Error(0)
}

func (vrrm *VoteRedisRepositoryMock) FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error) {
	args := vrrm.Called(ctx, userID)
	return args.Get(0).([]*model.UserVote), args.Error(1)
}

func (vrrm *VoteRedisRepositoryMock) SetFindUserVoteByUserID(ctx context.Context, userID *uuid.UUID, userVotes []*model.UserVote) error {
	args := vrrm.Called(ctx, userID, userVotes)
	return args.Error(0)
}

func (vrrm *VoteRedisRepositoryMock) FindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID) (*model.UserVote, error) {
	args := vrrm.Called(ctx, userVoteID)
	return args.Get(0).(*model.UserVote), args.Error(1)
}

func (vrrm *VoteRedisRepositoryMock) SetFindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID, userVote *model.UserVote) error {
	args := vrrm.Called(ctx, userVoteID, userVote)
	return args.Error(0)
}
