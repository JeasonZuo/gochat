package main

import (
	"fmt"
	"gochat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UsersModel{})

	db.Create(&models.UsersModel{Identity: "10000", Name: "zjx", Password: "123"})

	fmt.Println(db)
}
