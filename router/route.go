package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/const/routes"
	"kk-rschian.com/redis_auth/controller"
)

// FIXME: ルートを外部の値にしておかないと、mail serviceがルートの値を使えない
// 循環参照
func InitRoute() *gin.Engine {
	router := gin.Default()
	// static files
	router.Static("/assets", "views/assets")
	router.LoadHTMLGlob("views/*html")

	router.POST(routes.SignUp, controller.SignUp)
	router.GET(routes.VerifyEmail, controller.VerifyUser)
	router.POST(routes.Login, controller.Login)
	router.POST(routes.Logout, controller.Logout)
	router.POST(routes.RequestPasswordReset, controller.RequestResetPassword)
	router.GET(routes.PasswordResetForm, controller.ResetPasswordForm)
	// update: update userdata (need session)
	return router
}
