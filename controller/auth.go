package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschian.com/redis_auth/service/database"
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
		Name:       newUserInfo.Name,
		Email:      newUserInfo.Email,
		Password:   hashedPassword,
		IsVerified: false,
	}

	if err := database.CreateNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
