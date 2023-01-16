package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/config"
	"kk-rschian.com/redis_auth/redis"
	"kk-rschian.com/redis_auth/router"
)

func init() {
	config.Setup()
	redis.Setup()
}

func main() {
	gin.SetMode(config.Server.RunMode)
	router := router.InitRoute()

	service := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.HttpPort),
		Handler:      router,
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
	}
	service.ListenAndServe()
}
