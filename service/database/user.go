package database

import (
	"errors"

	"kk-rschian.com/redis_auth/const/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) (*models.User, error) {
	var user *models.User
	if err := DB.Where("ID=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	result := DB.Where("id=?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return nil
}

func UpdateUser(user *models.User) error {
	result := DB.Model(&user).Updates(map[string]interface{}{
		"name":     user.Name,
		"password": user.Password,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return nil
}
