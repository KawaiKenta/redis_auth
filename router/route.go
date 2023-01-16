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
		redis.SetSession(c, "hello", "world")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/get", func(c *gin.Context) {
		val, err := redis.GetSession(c, "hello")
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"key": val,
		})
	})

	router.GET("/del", func(c *gin.Context) {
		if err := redis.DeleteSession(c, "hello"); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"delete": "ok",
		})
	})
	return router
}
