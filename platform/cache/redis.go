package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func getRedisConnectionString() string {
	var connectionString string
	user := os.Getenv("REDIS_USER")
	password := os.Getenv("REDIS_PASSWORD")
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	if (user != "") && (password != "") {
		connectionString = fmt.Sprintf("redis://%s:%s@%s:%s", user, password, host, port)
	} else {
		connectionString = fmt.Sprintf("redis://%s:%s", host, port)
	}
	dbNumber := os.Getenv("REDIS_DB_NUMBER")
	if dbNumber != "" {
		connectionString += fmt.Sprintf("/%s", dbNumber)
	}
	return connectionString
}

func RedisConnect() (*redis.Client, error) {
	connectionString := getRedisConnectionString()
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(opt)
	return redisClient, nil
}
