package repository

import (
	"context"
	"encoding/json"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/infrastructure/datastore"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	votePrefix     = "vote:"
	userVotePrefix = "user_vote:"
	voteTtl        = 1 * time.Minute
)

type VoteRedisRepository interface {
	FindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID) ([]*model.Vote, error)
	SetFindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID, votes []*model.Vote) error
	FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error)
	SetFindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID, votes []*model.Vote) error
	FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error)
	SetFindUserVoteByUserID(ctx context.Context, userID *uuid.UUID, userVotes []*model.UserVote) error
	FindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID) (*model.UserVote, error)
	SetFindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID, userVote *model.UserVote) error
}

type voteRedisRepo struct {
	redis *datastore.Redis
}

func NewVoteRedisRepository(redis *datastore.Redis) VoteRedisRepository {
	return &voteRedisRepo{redis: redis}
}

func (vr *voteRedisRepo) FindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID) ([]*model.Vote, error) {
	key := vr.makeKey(votePrefix, voteUserID.String())
	votesBytes, err := vr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.VoteRedisRepoFindVoteByUserIDGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRedisRepoFindVoteByUserIDGet.AppendMessage(err)
	}

	votes := make([]*model.Vote, 0)
	err = json.Unmarshal(votesBytes, &votes)
	if err != nil {
		return nil, apperrors.VoteRedisRepoFindVoteByUserIDUnmarshal.AppendMessage(err)
	}
	return votes, nil
}

func (vr *voteRedisRepo) SetFindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID, votes []*model.Vote) error {
	key := vr.makeKey(votePrefix, voteUserID.String())
	votesBytes, err := json.Marshal(votes)
	if err != nil {
		return apperrors.VoteRedisRepoSetFindVoteByUserIDMarshal.AppendMessage(err)
	}

	err = vr.redis.RedisClient.Set(ctx, key, votesBytes, voteTtl).Err()
	if err != nil {
		return apperrors.VoteRedisRepoSetFindVoteByUserIDSet.AppendMessage(err)
	}
	return nil
}

func (vr *voteRedisRepo) FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error) {
	preparedKey, err := vr.prepareKey(userIDs)
	if err != nil {
		return nil, apperrors.VoteRedisRepoFindVotesByUserIDsPrepareKey.AppendMessage(err)
	}

	key := vr.makeKey(votePrefix, preparedKey)
	votesBytes, err := vr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.VoteRedisRepoFindVotesByUserIDsGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRedisRepoFindVotesByUserIDsGet.AppendMessage(err)
	}

	votes := make([]*model.Vote, 0)
	err = json.Unmarshal(votesBytes, &votes)
	if err != nil {
		return nil, apperrors.VoteRedisRepoFindVotesByUserIDsUnmarshal.AppendMessage(err)
	}
	return votes, nil
}

func (vr *voteRedisRepo) SetFindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID, votes []*model.Vote) error {
	preparedKey, err := vr.prepareKey(userIDs)
	if err != nil {
		return apperrors.VoteRedisRepoSetFindVotesByUserIDsPrepareKey.AppendMessage(err)
	}

	key := vr.makeKey(votePrefix, preparedKey)
	votesBytes, err := json.Marshal(votes)
	if err != nil {
		return apperrors.VoteRedisRepoSetFindVotesByUserIDsMarshal.AppendMessage(err)
	}

	err = vr.redis.RedisClient.Set(ctx, key, votesBytes, voteTtl).Err()
	if err != nil {
		return apperrors.VoteRedisRepoSetFindVotesByUserIDsSet.AppendMessage(err)
	}
	return nil
}

func (vr *voteRedisRepo) FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error) {
	key := vr.makeKey(userVotePrefix, userID.String())
	userVotesBytes, err := vr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.VoteRedisRepoFindUserVoteByUserIDGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRedisRepoFindUserVoteByUserIDGet.AppendMessage(err)
	}

	userVotes := make([]*model.UserVote, 0)
	err = json.Unmarshal(userVotesBytes, &userVotes)
	if err != nil {
		return nil, apperrors.VoteRedisRepoFindUserVoteByUserIDUnmarshal.AppendMessage(err)
	}
	return userVotes, nil
}

func (vr *voteRedisRepo) SetFindUserVoteByUserID(ctx context.Context, userID *uuid.UUID, userVotes []*model.UserVote) error {
	key := vr.makeKey(userVotePrefix, userID.String())
	userVotesBytes, err := json.Marshal(userVotes)
	if err != nil {
		return apperrors.VoteRedisRepoSetFindUserVoteByUserIDMarshal.AppendMessage(err)
	}

	err = vr.redis.RedisClient.Set(ctx, key, userVotesBytes, voteTtl).Err()
	if err != nil {
		return apperrors.VoteRedisRepoSetFindUserVoteByUserIDSet.AppendMessage(err)
	}
	return nil
}

func (vr *voteRedisRepo) FindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID) (*model.UserVote, error) {
	key := vr.makeKey(userVotePrefix, userVoteID.String())
	userVoteBytes, err := vr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.VoteRedisRepoFindUserVoteByIDGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRedisRepoFindUserVoteByIDGet.AppendMessage(err)
	}

	userVote := &model.UserVote{}
	err = json.Unmarshal(userVoteBytes, userVote)
	if err != nil {
		return nil, apperrors.VoteRedisRepoFindUserVoteByIDUnmarshal.AppendMessage(err)
	}
	return userVote, nil
}

func (vr *voteRedisRepo) SetFindUserVoteByID(ctx context.Context, userVoteID *uuid.UUID, userVote *model.UserVote) error {
	key := vr.makeKey(userVotePrefix, userVoteID.String())
	userVoteBytes, err := json.Marshal(userVote)
	if err != nil {
		return apperrors.VoteRedisRepoSetFindUserVoteByIDMarshal.AppendMessage(err)
	}

	err = vr.redis.RedisClient.Set(ctx, key, userVoteBytes, voteTtl).Err()
	if err != nil {
		return apperrors.VoteRedisRepoSetFindUserVoteByIDSet.AppendMessage(err)
	}
	return nil
}

func (vr *voteRedisRepo) prepareKey(data any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (vr *voteRedisRepo) makeKey(prefix string, key string) string {
	return prefix + key
}
