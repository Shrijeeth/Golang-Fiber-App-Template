package configs

import (
	"context"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/platform/cache"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() error {
	var err error

	RedisClient, err = cache.RedisConnect()
	if err != nil {
		return err
	}

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}

func CloseRedisClient() error {
	if RedisClient == nil {
		return nil
	}

	err := RedisClient.Close()
	if err != nil {
		return err
	}

	return nil
}
