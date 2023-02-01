package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
IsVerified: Emailが認証されているか
VerifyToken: Email認証, Password再設定時のアクセストークン
*/
type User struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name,omitempty"`
	Email    string `gorm:"size:255;not null;unique" json:"email,omitempty"`
	Password string `gorm:"size:255;not null" json:"password,omitempty"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
