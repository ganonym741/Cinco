package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"gitlab.com/cinco/configs"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.Config().RedisConfig.Host,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("can't connect to redis")
	}

	fmt.Println("Connected to redis")
	return client
}
