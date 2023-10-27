package main

import (
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/pkg/utils"
	"github.com/JeasonZuo/gochat/routers"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//config.Setup()
	models.Setup()
	utils.Setup()
}

func main() {
	router := routers.InitApiRouter()
	router.Run(":8080")
}
