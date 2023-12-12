package repository

import (
	"context"
	"encoding/json"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/infrastructure/datastore"
	"usermanager/internal/utils"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	userPrefix = "user:"
	userTtl    = 1 * time.Minute
)

type UserRedisRepository interface {
	FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	SetFindUserByUUID(ctx context.Context, userID uuid.UUID, user *model.User) error
	FindUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	SetFindUserByNickname(ctx context.Context, nickname string, user *model.User) error
	GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error)
	SetGetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery, users *model.Users) error
}

type userRedisRepo struct {
	redis *datastore.Redis
}

func NewUserRedisRepository(redis *datastore.Redis) UserRedisRepository {
	return &userRedisRepo{redis: redis}
}

func (ur *userRedisRepo) FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	key := ur.makeKey(userID.String())
	userBytes, err := ur.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.UserRedisRepoFindUserByUUIDGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.UserRedisRepoFindUserByUUIDGet.AppendMessage(err)
	}

	user := &model.User{}
	err = json.Unmarshal(userBytes, user)
	if err != nil {
		return nil, apperrors.UserRedisRepoFindUserByUUIDUnmarshal.AppendMessage(err)
	}

	return user, nil
}

func (ur *userRedisRepo) SetFindUserByUUID(ctx context.Context, userID uuid.UUID, user *model.User) error {
	key := ur.makeKey(userID.String())
	userBytes, err := json.Marshal(user)
	if err != nil {
		return apperrors.UserRedisRepoSetFindUserByUUIDMarshal.AppendMessage(err)
	}

	err = ur.redis.RedisClient.Set(ctx, key, userBytes, userTtl).Err()
	if err != nil {
		return apperrors.UserRedisRepoSetFindUserByUUIDSet.AppendMessage(err)
	}

	return nil
}

func (ur *userRedisRepo) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	key := ur.makeKey(nickname)
	userBytes, err := ur.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.UserRedisRepoFindUserByNicknameGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.UserRedisRepoFindUserByNicknameGet.AppendMessage(err)
	}

	user := &model.User{}
	err = json.Unmarshal(userBytes, user)
	if err != nil {
		return nil, apperrors.UserRedisRepoFindUserByNicknameUnmarshal.AppendMessage(err)
	}

	return user, nil
}

func (ur *userRedisRepo) SetFindUserByNickname(ctx context.Context, nickname string, user *model.User) error {
	key := ur.makeKey(nickname)
	userBytes, err := json.Marshal(user)
	if err != nil {
		return apperrors.UserRedisRepoSetFindUserByNicknameMarshal.AppendMessage(err)
	}

	err = ur.redis.RedisClient.Set(ctx, key, userBytes, userTtl).Err()
	if err != nil {
		return apperrors.UserRedisRepoSetFindUserByNicknameSet.AppendMessage(err)
	}

	return nil
}

func (ur *userRedisRepo) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	preparedKey, err := ur.prepareKey(paginationQuery)
	if err != nil {
		return nil, apperrors.UserRedisRepoGetUsersPrepareKey.AppendMessage(err)
	}
	key := ur.makeKey(preparedKey)
	usersBytes, err := ur.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.UserRedisRepoGetUsersGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.UserRedisRepoGetUsersGet.AppendMessage(err)
	}

	users := &model.Users{}
	err = json.Unmarshal(usersBytes, users)
	if err != nil {
		return nil, apperrors.UserRedisRepoGetUsersUnmarshal.AppendMessage(err)
	}

	return users, nil
}

func (ur *userRedisRepo) SetGetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery, users *model.Users) error {
	preparedKey, err := ur.prepareKey(paginationQuery)
	if err != nil {
		return apperrors.UserRedisRepoSetGetUsersPrepareKey.AppendMessage(err)
	}
	key := ur.makeKey(preparedKey)
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return apperrors.UserRedisRepoSetGetUsersMarshal.AppendMessage(err)
	}

	err = ur.redis.RedisClient.Set(ctx, key, usersBytes, userTtl).Err()
	if err != nil {
		return apperrors.UserRedisRepoSetGetUsersSet.AppendMessage(err)
	}

	return nil
}

func (ur *userRedisRepo) prepareKey(data any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (ur *userRedisRepo) makeKey(key string) string {
	return userPrefix + key
}
