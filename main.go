package main

import (
	"github.com/JeasonZuo/gochat/config"
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/routers"
	"github.com/JeasonZuo/gochat/utils"
)

func init() {
	config.Setup()
	models.Setup()
	utils.Setup()
}

func main() {
	router := routers.InitApiRouter()
	router.Run(":8001")
}
