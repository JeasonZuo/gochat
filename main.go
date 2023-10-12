package main

import (
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/pkg/config"
	"github.com/JeasonZuo/gochat/pkg/gredis"
	"github.com/JeasonZuo/gochat/pkg/utils"
	"github.com/JeasonZuo/gochat/routers"
)

func init() {
	config.Setup()
	models.Setup()
	gredis.SetUp()
	utils.Setup()
}

func main() {
	router := routers.InitApiRouter()
	router.Run(":8001")
}
