package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
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
	// TODO: return message
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(newUserInfo.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 新規ユーザー作成
	user := database.User{
		Name:     newUserInfo.Name,
		Email:    newUserInfo.Email,
		Password: hashedPassword,
	}

	if err := database.CreateNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Email認証を行い、DBにユーザーを登録する
// Emailからのアクセスなのでページを返してあげる
func VerifyUser(c *gin.Context) {
	// get access token from url
	token := c.Query("uuid")

	// check existance from redis
	userJson, err := redis.GetSession(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	var user database.User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}

	// create user
	if err := database.CreateNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// TODO: jsonではなくページを返す必要があるのでは
	c.JSON(http.StatusOK, user)
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

func TestView(c *gin.Context) {
	// get access token from url
	token := c.Query("uuid")
	println(token)
	c.HTML(http.StatusOK, "verify.html", gin.H{"token": token})
}
