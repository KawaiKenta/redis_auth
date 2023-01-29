package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
	"kk-rschian.com/redis_auth/service/redis"
	"kk-rschian.com/redis_auth/utils"
)

// Email認証を行い、DBにユーザーを登録する
// Emailからのアクセスなのでページを返してあげる
func VerifyUser(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	// check existance from redis
	userJson, err := redis.GetUserInfo(c, uuid)
	if err != nil {
		// 認証情報がありません
		c.HTML(http.StatusBadRequest, "verify_failed_expire.html", gin.H{"token": uuid})
		return
	}

	var user database.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		// ユーザーデータのパース中にエラー
		c.HTML(http.StatusInternalServerError, "verify_failed_internal.html", gin.H{"token": uuid})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		// ハッシュ化に失敗
		c.HTML(http.StatusInternalServerError, "verify_failed_internal.html", gin.H{"token": uuid})
		return
	}
	user.Password = hashedPassword

	// create user
	if err := database.CreateNewUser(&user); err != nil {
		// データベースサーバーにエラー
		c.HTML(http.StatusBadRequest, "verify_failed_deplicate.html", gin.H{"token": uuid})
		return
	}

	redis.DeleteUserInfo(c, uuid)
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
	// userJson, err := redis.GetUserInfo(c, uuid)
	// if err != nil {
	// 	// 認証情報がありません
	// 	c.HTML(http.StatusBadRequest, "verify_failed_expire.html", gin.H{"token": uuid})
	// 	return
	// }
}
