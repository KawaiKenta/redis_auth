package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/redis"
)

func SetSession(c *gin.Context) {
	if err := redis.SetSession(c, "heler23lo", "worl234d"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetSession(c *gin.Context) {
	val, err := redis.GetSession(c, "hello")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"key": val,
	})
}

func DeleteSession(c *gin.Context) {
	if err := redis.DeleteSession(c, "hello"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "ok",
	})
}

// 	_, err := c.Cookie(key)
// c.SetCookie(key, "", -1, "/", "localhost", secure, httpOnly)
// c.SetCookie(key, value, 0, "/", "localhost", secure, httpOnly)
// 	_, err := c.Cookie(key)
// if err != nil {
// 	return "", err
// }
