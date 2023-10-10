package main

import (
	"gochat/routers"
	"gochat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	router := routers.InitApiRouter()
	router.Run()
}
