package main

import (
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/pkg/config"
	"github.com/JeasonZuo/gochat/pkg/gredis"
	"github.com/JeasonZuo/gochat/pkg/utils"
	"github.com/JeasonZuo/gochat/routers"
	"github.com/JeasonZuo/gochat/service/ws_service"
)

func init() {
	config.Setup()
	models.Setup()
	gredis.SetUp()
	utils.Setup()
	ws_service.SetUp()
}

func main() {
	router := routers.InitApiRouter()
	router.Run(":8001")
}
