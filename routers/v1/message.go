package v1

import (
	"github.com/JeasonZuo/gochat/pkg/app"
	"github.com/JeasonZuo/gochat/service/message_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMessageList(c *gin.Context) {
	appG := app.Gin{C: c}
	strToUserId := c.Query("to_user_id")

	toUserId, _ := strconv.ParseUint(strToUserId, 10, 64)
	message := message_service.Message{
		ToUserId:   uint(toUserId),
		FromUserId: c.GetUint("loginUserId"),
	}

	messageList, err := message.GetMessageList()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 10002, err.Error(), nil)
	}

	appG.Response(http.StatusOK, 10000, "ok", messageList)
}
