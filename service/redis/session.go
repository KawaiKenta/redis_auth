package redis

import (
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, sessionId string) error {
	if err := Client.Set(c, "sessionId", sessionId, expTime).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteSession(c *gin.Context, key string) error {
	if err := Client.Del(c, key).Err(); err != nil {
		return err
	}
	return nil
}

func GetSession(c *gin.Context, key string) (string, error) {
	val, err := Client.Get(c, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
