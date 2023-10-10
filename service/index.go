package service

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary 测试连通性
// @Success 200 {string} Go chat
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Go chat",
	})
}
