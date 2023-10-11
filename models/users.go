package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name      string `json:"name" gorm:"size:31"`
	AvatarUrl string `json:"avatar_url" gorm:"size:255"`
	Password  string `json:"password" gorm:"size:63"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func GetUserList() []*UserModel {
	data := make([]*UserModel, 10)
	db.Find(&data)
	return data
}

func AddUser(data map[string]any) (uint, error) {
	userModel := UserModel{
		Name:      data["name"].(string),
		AvatarUrl: data["avatar_url"].(string),
		Password:  data["password"].(string),
	}
	if err := db.Create(&userModel).Error; err != nil {
		return 0, err
	}

	return userModel.ID, nil
}
