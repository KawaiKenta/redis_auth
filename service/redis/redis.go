package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
	"kk-rschian.com/redis_auth/config"
)

var (
	Client  *redis.Client
	expTime time.Duration
)

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DataBaseType,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis server error: %v", err)
	}
}

func Close() {
	if err := Client.Close(); err != nil {
		log.Fatalf("redis server error: %v", err)
	}
}
