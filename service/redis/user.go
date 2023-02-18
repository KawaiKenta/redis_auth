package redis

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
)

func GetUserInfo(c *gin.Context, key string) (*database.User, error) {
	// redis からの取得
	userJson, err := Client.Get(c, key).Result()
	if err != nil {
		return nil, err
	}

	// userへのパース
	var user *database.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserInfo(c *gin.Context, key string) error {
	if err := Client.Del(c, key).Err(); err != nil {
		return err
	}
	return nil
}

func SetUserInfo(c *gin.Context, key string, user *database.User) error {
	serialize, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := Client.Set(c, key, serialize, time.Minute*30).Err(); err != nil {
		return err
	}
	return nil
}
