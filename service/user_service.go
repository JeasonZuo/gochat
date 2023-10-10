package service

import (
	"github.com/gin-gonic/gin"
	"gochat/models"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UsersModel, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"data": data,
	})
}
