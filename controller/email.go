package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
	"kk-rschian.com/redis_auth/service/redis"
)

// Email認証を行い、DBにユーザーを登録する
// Emailからのアクセスなのでページを返してあげる
func VerifyUser(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	// check existance from redis
	user, err := redis.GetUser(c, uuid)
	if err != nil {
		// 認証情報がありません
		c.HTML(http.StatusBadRequest, "verify_failed_expire.html", gin.H{"token": uuid})
		return
	}

	// パスワードのハッシュ化
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// create user
	if err := database.CreateUser(user); err != nil {
		// データベースサーバーにエラー
		c.HTML(http.StatusBadRequest, "verify_failed_deplicate.html", gin.H{"token": uuid})
		return
	}

	redis.DeleteUser(c, uuid)
	c.HTML(http.StatusOK, "verify_success.html", gin.H{"token": uuid})
}

// 2. show reset form
// get new password
// send it with uuid
// check redis
// overwrite
// show page
func ResetPasswordForm(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	_, err := redis.GetUser(c, uuid)
	if err != nil {
		// 認証情報がありません
		c.HTML(http.StatusBadRequest, "verify_failed_expire.html", gin.H{"token": uuid})
		return
	}

	c.HTML(http.StatusOK, "verify_success.html", gin.H{"token": uuid})

	// userJson, err := redis.GetUserInfo(c, uuid)
	// if err != nil {
	// 	// 認証情報がありません
	// 	c.HTML(http.StatusBadRequest, "verify_failed_expire.html", gin.H{"token": uuid})
	// 	return
	// }
}
