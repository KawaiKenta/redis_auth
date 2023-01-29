package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/controller"
)

// FIXME: ルートを外部の値にしておかないと、mail serviceがルートの値を使えない
// 循環参照
func InitRoute() *gin.Engine {
	router := gin.Default()
	// static files
	router.Static("/assets", "views/assets")
	router.LoadHTMLGlob("views/*html")

	router.POST("/user/signup", controller.SignUp)
	router.GET("/user/verify/:uuid", controller.VerifyUser)
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", controller.Logout)
	router.POST("/user/request-password-reset", controller.RequestResetPassword)
	router.GET("/user/reset-password-form/:uuid", controller.ResetPasswordForm)
	// update: update userdata (need session)
	return router
}
