package models

import (
	"gochat/utils"
	"gorm.io/gorm"
)

type UsersModel struct {
	gorm.Model
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UsersModel) TableName() string {
	return "tb_users"
}

func GetUserList() []*UsersModel {
	data := make([]*UsersModel, 10)
	utils.DB.Find(&data)
	return data
}
