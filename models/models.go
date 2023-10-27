package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func Setup() {
	dsn := os.Getenv("MYSQL_DSN")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dbSession, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db = dbSession

	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FriendModel{})
	db.AutoMigrate(&MessageModel{})
}
