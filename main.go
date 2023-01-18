package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/config"
	"kk-rschian.com/redis_auth/router"
	"kk-rschian.com/redis_auth/service/database"
	"kk-rschian.com/redis_auth/service/redis"
	"kk-rschian.com/redis_auth/utils"
)

func init() {
	config.Setup()
	redis.Setup()
	database.Setup()
	for i := 0; i < 100; i++ {
		println(utils.CreateToken())
	}
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
