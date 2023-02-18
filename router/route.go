package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/const/routes"
	"kk-rschian.com/redis_auth/controller"
	"kk-rschian.com/redis_auth/middleware"
)

// FIXME: ルートを外部の値にしておかないと、mail serviceがルートの値を使えない
// 循環参照
func InitRoute() *gin.Engine {
	router := gin.Default()
	user := router.Group("/user", middleware.AddCorsHeader)
	{
		user.POST(routes.SignUp, controller.Signup)
		user.GET(routes.VerifyEmail, controller.VerifyEmail)
		user.POST(routes.Login, controller.Login)
		user.POST(routes.Logout, controller.Logout)
		user.POST(routes.ForgetPassword, controller.RequestResetPassword)
		user.POST(routes.ResetPassword, controller.ResetPassword)
	}

	return router
}
