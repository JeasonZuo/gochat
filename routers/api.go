package routers

import (
	"github.com/JeasonZuo/gochat/docs"
	cros "github.com/JeasonZuo/gochat/middleware/cors"
	"github.com/JeasonZuo/gochat/middleware/jwt"
	v1 "github.com/JeasonZuo/gochat/routers/v1"
	"github.com/JeasonZuo/gochat/service"
	"github.com/JeasonZuo/gochat/service/ws_service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitApiRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cros.Cors())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/index", service.GetIndex)
		apiV1.POST("/sign_up", v1.UserRegister) //用户注册
		apiV1.POST("/sign_in", v1.UserLogin)    //用户登录

		apiV1.Use(jwt.JWT())
		{
			apiV1.GET("/user_info", v1.GetUserInfo) //获取用户信息
			//编辑用户信息
			apiV1.POST("/add_friend", v1.AddFriend)         //添加好友
			apiV1.GET("/get_friend_list", v1.GetFriendList) //获取好友列表
			//删除好友

			apiV1.GET("/get_message_list", v1.GetMessageList) //获取历史信息
		}
	}

	r.GET("/ws", ws_service.WebSocketHandler)

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
