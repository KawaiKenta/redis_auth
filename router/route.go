package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/redis"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/set", func(c *gin.Context) {
		err := redis.Redis.Set(c, "testkey", "hello", 0).Err()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/get", func(c *gin.Context) {
		val, err := redis.Redis.Get(c, "testkey").Result()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"key": val,
		})
	})
	return router
}
