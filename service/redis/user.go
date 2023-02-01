package redis

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/models"
)

func SetUser(c *gin.Context, key string, user *models.User) error {
	serialize, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := Client.Set(c, key, serialize, time.Minute*30).Err(); err != nil {
		return err
	}
	return nil
}

func GetUser(c *gin.Context, key string) (*models.User, error) {
	var user *models.User
	serialize, err := Client.Get(c, key).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(serialize), &user); err != nil {
		// ユーザーデータのパース中にエラー
		return nil, err
	}
	return user, nil
}

func DeleteUser(c *gin.Context, key string) error {
	if err := Client.Del(c, key).Err(); err != nil {
		return err
	}
	return nil
}
