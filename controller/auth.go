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

// Email認証を行い、DBにユーザーを登録する
// Emailからのアクセスなのでページを返してあげる
func VerifyUser(c *gin.Context) {
	// get access token from url
	uuid := c.Query("uuid")

	// check existance from redis
	userJson, err := redis.GetUserInfo(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	var user database.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}

	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	user.Password = hashedPassword

	// create user
	if err := database.CreateNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	redis.DeleteUserInfo(c, uuid)
	c.HTML(http.StatusOK, "verify.html", gin.H{"token": uuid})
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
