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

	// 2回目だとエラーが帰る, redisのセッションが切れる？
	router.POST("/user/signup", controller.SignUp)
	router.GET("/user/verify", controller.VerifyUser)

	router.GET("/set", controller.SetSession)
	router.GET("/get", controller.GetSession)
	router.GET("/del", controller.DeleteSession)

	router.GET("/test", controller.TestMail)
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
