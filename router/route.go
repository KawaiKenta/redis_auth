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
	router.GET("/user/verify", controller.VerifyUser)
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", controller.Logout)

	// password reset
	// 1. send reset mail
	// get email
	// create uuid
	// set session
	// send email

	// 2. show reset form
	// get new password
	//

	// update: update userdata (need session)
	return router
}
