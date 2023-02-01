package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// すでにユーザーが存在している場合はエラー
	user, _ := database.GetUserByEmail(newUserInfo.Email)
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "すでに使用されているメールアドレスです"})
		return
	}

	// userをredisに登録
	uuid := utils.CreateToken()
	if err := redis.SetUser(c, uuid, user); err != nil {
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

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// データベースにuserが存在するかチェック
	user, err := database.GetUserByEmail(loginForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "使用されていないメールアドレスです"})
		return
	}

	// パスワードを検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "パスワードが間違っています"})
		return
	}

	// セッションサーバーへ送る
	sessionId := utils.CreateToken()
	if err := redis.SetUser(c, sessionId, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// cookieのセット
	c.SetCookie("sessionId", sessionId, 0, "/", "localhost", false, false)
	c.Status(http.StatusNoContent)
}

func Logout(c *gin.Context) {
	// cookieの取得
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// セッションサーバーから消去
	if err := redis.DeleteUser(c, sessionId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// cookieのオーバーライト
	c.SetCookie("sessionId", "", -1, "/", "localhost", false, false)
	c.Status(http.StatusNoContent)
}

func RequestResetPassword(c *gin.Context) {
	// ユーザー登録情報
	var requestForm struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&requestForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// データベースにuserが存在するかチェック
	user, err := database.GetUserByEmail(requestForm.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "使用されていないメールアドレスです"})
		return
	}

	// uuidをkeyとして、redisに保存
	uuid := utils.CreateToken()
	if err := redis.SetUser(c, uuid, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 認証用メールの送信
	if err := mail.SendResetPasswordMail(c, requestForm.Email, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
