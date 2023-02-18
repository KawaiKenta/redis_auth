package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/controller"
	"kk-rschian.com/redis_auth/middleware"
)

// FIXME: ルートを外部の値にしておかないと、mail serviceがルートの値を使えない
// 循環参照
func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AddCorsHeader)

	router.POST("/user/signup", controller.Signup)
	router.GET("/user/verify", controller.VerifyEmail)
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", controller.Logout)
	router.POST("/user/forgetpassword", controller.RequestResetPassword)
	router.POST("/user/resetpassword", controller.ResetPassword)
	return router
}
