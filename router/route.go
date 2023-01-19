package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/controller"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	// static files
	router.Static("/assets", "views/assets")
	router.LoadHTMLGlob("views/*html")

	// routing
	router.GET("/set", controller.SetSession)
	// router.GET("/get", controller.GetSession)
	router.GET("/del", controller.DeleteSession)
	router.GET("/verify", controller.TestView)
	// signup: set userdata to db. set random id for verify.
	//         then, send email.
	// verify: verify email. use db
	//          you have to retry signup
	// サインアップ時にデータベースを使うかredisを使うか悩みどころ
	// login: get login data and verify with database. After that,
	//        set the session data, then return session data.
	// logout: delete session data
	// update: update userdata (need session)
	return router
}
