package usecase

import (
	"context"

	"usermanager/internal/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type VoteRepositoryMock struct {
	mock.Mock
}

func (vrm *VoteRepositoryMock) SaveVote(ctx context.Context, vote *model.Vote) (*model.Vote, error) {
	args := vrm.Called(ctx, vote)
	return args.Get(0).(*model.Vote), args.Error(1)
}

func (vrm *VoteRepositoryMock) UpdateVote(ctx context.Context, vote *model.Vote) (*model.Vote, error) {
	args := vrm.Called(ctx, vote)
	return args.Get(0).(*model.Vote), args.Error(1)
}

func (vrm *VoteRepositoryMock) FindVoteByID(ctx context.Context, voteID int64) (*model.Vote, error) {
	args := vrm.Called(ctx, voteID)
	return args.Get(0).(*model.Vote), args.Error(1)
}

func (vrm *VoteRepositoryMock) FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error) {
	args := vrm.Called(ctx, userID)
	return args.Get(0).([]*model.UserVote), args.Error(1)
}

func (vrm *VoteRepositoryMock) FindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID) ([]*model.Vote, error) {
	args := vrm.Called(ctx, voteUserID)
	return args.Get(0).([]*model.Vote), args.Error(1)
}

func (vrm *VoteRepositoryMock) FindUserVoteByID(ctx context.Context, userVoteID int64) (*model.UserVote, error) {
	args := vrm.Called(ctx, userVoteID)
	return args.Get(0).(*model.UserVote), args.Error(1)
}

func (vrm *VoteRepositoryMock) SaveUserVote(ctx context.Context, userVote *model.UserVote) (*model.UserVote, error) {
	args := vrm.Called(ctx, userVote)
	return args.Get(0).(*model.UserVote), args.Error(1)
}

func (vrm *VoteRepositoryMock) DeleteUserVote(ctx context.Context, userVote *model.UserVote) error {
	args := vrm.Called(ctx, userVote)
	return args.Error(1)
}

func (vrm *VoteRepositoryMock) DeleteVote(ctx context.Context, vote *model.Vote) error {
	args := vrm.Called(ctx, vote)
	return args.Error(1)
}

func (vrm *VoteRepositoryMock) FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error) {
	args := vrm.Called(ctx, userIDs)
	return args.Get(0).([]*model.Vote), args.Error(1)
}
