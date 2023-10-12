package jwt

import (
	"github.com/JeasonZuo/gochat/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = 0
		var msg = ""

		token := c.Query("token")
		if token == "" {
			code = 10001
			msg = "token is required"
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					code = 10001
					msg = "token is expired"
				default:
					code = 10001
					msg = "token is invalid"
				}
			}
			c.Set("loginUserId", claims.ID)
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
				"data": nil,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
