package redis

import (
	"github.com/gin-gonic/gin"
)

// TODO: sessionのhttponlyとかは有効にしたほうがいい
func SetSession(c *gin.Context, key string, value string) error {
	if err := Client.Set(c, key, value, expTime).Err(); err != nil {
		return err
	}
	c.SetCookie(key, value, 0, "/", "localhost", secure, httpOnly)
	return nil
}

func DeleteSession(c *gin.Context, key string) error {
	_, err := c.Cookie(key)
	if err != nil {
		return err
	}
	if err := Client.Del(c, key).Err(); err != nil {
		return err
	}
	c.SetCookie(key, "", -1, "/", "localhost", secure, httpOnly)
	return nil
}

func GetSession(c *gin.Context, key string) (string, error) {
	_, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	val, err := Client.Get(c, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
