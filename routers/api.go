package routers

import (
	"github.com/JeasonZuo/gochat/docs"
	v1 "github.com/JeasonZuo/gochat/routers/v1"
	"github.com/JeasonZuo/gochat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitApiRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/index", service.GetIndex)

		userGroup := apiV1.Group("/user")
		{
			userGroup.POST("/sign_up", v1.RegisterUser)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "route not found",
			"data":    nil,
		})
	})

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
