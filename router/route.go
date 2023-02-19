package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/const/routes"
	"kk-rschian.com/redis_auth/controller"
	"kk-rschian.com/redis_auth/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AddCorsHeader)
	user := router.Group("/user")
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
