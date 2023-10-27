package main

import (
	"fmt"
	"github.com/JeasonZuo/gochat/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.FriendModel{})
	db.AutoMigrate(&models.MessageModel{})

	fmt.Println(db)
}
