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
		c.HTML(http.StatusOK, "verify_failed.html", gin.H{"token": uuid})
		return
	}

	var user database.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		// ユーザーデータのパース中にエラー
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		// ハッシュ化に失敗
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	user.Password = hashedPassword

	// create user
	if err := database.CreateNewUser(&user); err != nil {
		// データベースサーバーにエラー
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	redis.DeleteUserInfo(c, uuid)
	c.HTML(http.StatusOK, "verify_success.html", gin.H{"token": uuid})
}
