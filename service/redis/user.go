package redis

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/const/models"
)

func GetUser(c *gin.Context, key string) (*models.User, error) {
	// redis からの取得
	userJson, err := Client.Get(c, key).Result()
	if err != nil {
		return nil, err
	}
	// userへのパース
	var user *models.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
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
