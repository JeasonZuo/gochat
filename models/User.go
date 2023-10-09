package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//func (table *User) TableName() string {
//	return "user"
//}
