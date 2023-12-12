package datastore

import (
	"context"
	"fmt"

	"usermanager/internal/apperrors"
	"usermanager/internal/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	RedisClient *redis.Client
}

const (
	MinIdleConns = 10
	PoolSize     = 10
	PoolTimeout  = 5
	redisDB      = 0
)

func NewRedisClient(cfg *config.Config) (*Redis, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.RedisHost, cfg.Redis.RedisPort),
		Username:     cfg.Redis.RedisUser,
		Password:     cfg.Redis.RedisPassword,
		MinIdleConns: MinIdleConns,
		PoolSize:     PoolSize,
		PoolTimeout:  PoolTimeout,
		DB:           redisDB,
	})

	err := pingRedis(ctx, client)
	if err != nil {
		return nil, apperrors.PingRedisError.AppendMessage(err)
	}

	return &Redis{RedisClient: client}, nil
}

func pingRedis(ctx context.Context, client *redis.Client) error {
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
