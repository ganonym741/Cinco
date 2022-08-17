package redis

import (
	"github.com/go-redis/redis/v9"
	"gitlab.com/cinco/configs"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.Config().RedisConfig.Host,
		Password: "",
		DB:       0,
	})
	return client
}
