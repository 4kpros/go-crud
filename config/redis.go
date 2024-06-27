package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectToRedis() (err error) {
	addr := fmt.Sprintf("%s:%d", AppEnvConfig.RedisHost, AppEnvConfig.RedisPort)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: AppEnvConfig.RedisUserName,
		Password: AppEnvConfig.RedisPassword,
		DB:       AppEnvConfig.RedisDatabase,
	})

	return
}

func SetRedisVal(key string, val string) (err error) {
	err = RedisClient.Set(ctx, key, val, 0).Err()
	return
}

func GetRegisVal(key string) (val string, err error) {
	val, err = RedisClient.Get(ctx, key).Result()
	return
}
