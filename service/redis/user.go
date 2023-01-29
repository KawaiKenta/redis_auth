package redis

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
)

func SetUserInfo(c *gin.Context, key string, json string) error {
	if err := Client.Set(c, key, json, time.Minute*30).Err(); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(c *gin.Context, key string) (string, error) {
	val, err := Client.Get(c, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func DeleteUserInfo(c *gin.Context, key string) error {
	if err := Client.Del(c, key).Err(); err != nil {
		return err
	}
	return nil
}

func SetResetPassword(c *gin.Context, key string, user *database.User) error {
	serialize, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := Client.Set(c, key, serialize, time.Minute*30).Err(); err != nil {
		return err
	}
	return nil
}
