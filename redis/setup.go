package redis

import (
	"fmt"
	"log"

	reds "github.com/go-redis/redis/v9"
	"kk-rschian.com/redis_auth/config"
)

var (
	Redis *reds.Client
)

func Setup() {
	Redis = reds.NewClient(&reds.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DataBaseType,
	})
	fmt.Println(config.Redis.Address)
	fmt.Println(config.Redis.Password)
	fmt.Println(config.Redis.DataBaseType)
}

func Close() {
	if err := Redis.Close(); err != nil {
		log.Fatalf("close when redis client: %v", err)
	}
}
