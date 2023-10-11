package main

import (
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/routers"
	"github.com/JeasonZuo/gochat/utils"
)

func init() {
	utils.InitConfig()
	models.Setup()
}

func main() {
	router := routers.InitApiRouter()
	router.Run(":8001")
}
