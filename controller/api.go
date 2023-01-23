package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
	"kk-rschian.com/redis_auth/service/mail"
	"kk-rschian.com/redis_auth/service/redis"
	"kk-rschian.com/redis_auth/utils"
)

func SignUp(c *gin.Context) {
	// ユーザー登録情報
	var newUserInfo struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	// veridation : エラーメッセージの改善
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// jsonデータに変換
	serialize, err := json.Marshal(&newUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// uuidをkeyとして、redisに保存
	uuid := utils.CreateToken()
	if err := redis.SetUserInfo(c, uuid, string(serialize)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 認証用メールの送信
	if err := mail.SendEmailVerifyMail(c, newUserInfo.Email, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func Login(c *gin.Context) {
	// ユーザー登録情報
	var loginForm struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	// veridation : エラーメッセージの改善
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// データベースにuserが存在するかチェック
	_, err := database.GetUserByEmail(loginForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// セッションサーバーへ送る
	sessionId := utils.CreateToken()
	if err := redis.SetSession(c, sessionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// cookieのセット
	c.SetCookie("sessionId", sessionId, 0, "/", "localhost", false, false)
	c.Status(http.StatusNoContent)
}

// passwordリセットの要求
func RequestPassword(c *gin.Context) {
	// リセット用uuidを作成
	// emailとともにreidsにセット
	// 贈りました
}

// passwordリセットを行う
// emailからなのでページを返してあげる
func ResetPassword(c *gin.Context) {

}

// 	_, err := c.Cookie(key)
// c.SetCookie(key, "", -1, "/", "localhost", secure, httpOnly)

// 	_, err := c.Cookie(key)
// if err != nil {
// 	return "", err
// }
