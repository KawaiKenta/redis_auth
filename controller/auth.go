package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"kk-rschian.com/redis_auth/const/models"
	"kk-rschian.com/redis_auth/const/validation"
	"kk-rschian.com/redis_auth/service/database"
	"kk-rschian.com/redis_auth/service/mail"
	"kk-rschian.com/redis_auth/service/redis"
	"kk-rschian.com/redis_auth/utils"
)

func Signup(c *gin.Context) {
	// ユーザー登録情報
	var form validation.UserRegister
	c.BindJSON(&form)
	if err := form.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// すでにユーザーが存在している場合はエラー
	user, _ := database.GetUserByEmail(form.Email)
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "すでに使用されているメールアドレスです"})
		return
	}

	// uuidをkeyとして、redisに保存
	user = &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	uuid := utils.CreateToken()
	if err := redis.SetUserInfo(c, uuid, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 認証用メールの送信
	if err := mail.SendEmailVerifyMail(c, form.Email, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ユーザー仮登録完了",
	})
}

func VerifyEmail(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	// check existance from redis
	user, err := redis.GetUserInfo(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なtokenです"})
		return
	}

	// パスワードのハッシュ化
	if err != user.HashPassword() {
		// ハッシュ化に失敗
		c.JSON(http.StatusInternalServerError, gin.H{"message": "パスワードのハッシュ化に失敗しました"})
		return
	}

	// create user
	if err := database.CreateUser(user); err != nil {
		// データベースサーバーにエラー
		c.JSON(http.StatusInternalServerError, gin.H{"message": "データベースエラー"})
		return
	}

	redis.DeleteUserInfo(c, uuid)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ユーザー本登録完了",
	})
}

func Login(c *gin.Context) {
	// ユーザー登録情報
	var form validation.Login
	c.BindJSON(&form)
	if err := form.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// データベースにuserが存在するかチェック
	user, err := database.GetUserByEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "使用されていないメールアドレスです"})
		return
	}

	// パスワードを検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "パスワードが間違っています"})
		return
	}

	// セッションサーバーへ送る
	sessionId := utils.CreateToken()
	if err := redis.SetUserInfo(c, sessionId, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// cookieのセット
	c.SetCookie("sessionId", sessionId, 0, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ログイン成功",
	})
}

func Logout(c *gin.Context) {
	// cookieの取得
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// セッションサーバーから消去
	if err := redis.DeleteUser(c, sessionId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// cookieのオーバーライト
	c.SetCookie("sessionId", "", -1, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ログアウト成功",
	})
}

func RequestResetPassword(c *gin.Context) {
	// ユーザー登録情報
	var form validation.ForgetPassword
	c.BindJSON(&form)
	if err := form.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// データベースにuserが存在するかチェック
	user, err := database.GetUserByEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "使用されていないメールアドレスです"})
		return
	}

	// uuidをkeyとして、redisに保存
	uuid := utils.CreateToken()
	if err := redis.SetUserInfo(c, uuid, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 認証用メールの送信
	if err := mail.SendResetPasswordMail(c, form.Email, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "パスワードリセット申請を受理しました",
	})
}

func ResetPassword(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	// パスワード情報
	var form validation.ResetPassword
	c.BindJSON(&form)
	if err := form.Validation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// ユーザー情報の取得
	user, err := redis.GetUserInfo(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なtokenです"})
		return
	}

	// パスワードのハッシュ化
	if err != user.HashPassword() {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "パスワードのハッシュ化に失敗しました"})
		return
	}

	// パスワードの更新
	if err := database.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "データベースエラー"})
		return
	}

	redis.DeleteUserInfo(c, uuid)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "パスワードを再設定しました",
	})
}
