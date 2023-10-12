package routers

import (
	"github.com/JeasonZuo/gochat/docs"
	"github.com/JeasonZuo/gochat/middleware/jwt"
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
			userGroup.POST("/sign_up", v1.UserRegister) //用户注册
			userGroup.POST("/sign_in", v1.UserLogin)    //用户登录
		}

		apiV1.Use(jwt.JWT())
		{
			//获取用户信息
			//编辑用户信息
			//添加好友
			//同意好友请求
			//获取好友列表
			//删除好友
			//创建群组
			//编辑群信息
			//加入群组
			//拉好友进群
			//退出群组
			//解散群组
			//获取群组列表
		}
	}

	//上传文件

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
