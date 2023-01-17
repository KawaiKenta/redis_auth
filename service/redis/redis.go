package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis/v9"
	"kk-rschian.com/redis_auth/config"
)

var (
	Client   *redis.Client
	expTime  time.Duration
	secure   bool
	httpOnly bool
)

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DataBaseType,
	})
	expTime = config.Redis.ExpirationTime
	// TODO: need change
	secure = false
	httpOnly = false
}

func Close() {
	if err := Client.Close(); err != nil {
		log.Fatalf("close when redis client: %v", err)
	}
}
