package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name      string `json:"name" gorm:"size:31"`
	AvatarUrl string `json:"avatar_url" gorm:"size:255"`
	Password  string `json:"password" gorm:"size:63"`
}

type UserInfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func AddUser(data map[string]any) (uint, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	passwordHashStr := string(hash)
	userModel := UserModel{
		Name:      data["name"].(string),
		AvatarUrl: data["avatar_url"].(string),
		Password:  passwordHashStr,
	}
	if err := db.Create(&userModel).Error; err != nil {
		return 0, err
	}

	return userModel.ID, nil
}

func EditUser(data map[string]any) error {
	userModel := UserModel{
		Model: gorm.Model{
			ID: data["id"].(uint),
		},
		Name:      data["name"].(string),
		AvatarUrl: data["avatar_url"].(string),
	}
	if err := db.Select("Name", "AvatarUrl").Save(&userModel).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id uint) (*UserModel, error) {
	userModel := &UserModel{}
	if err := db.First(userModel, id).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}
