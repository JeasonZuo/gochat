package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gochat/docs"
	"gochat/service"
)

func InitApiRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/index", service.GetIndex)

		userGroup := apiV1.Group("/user")
		{
			userGroup.GET("list", service.GetUserList)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
