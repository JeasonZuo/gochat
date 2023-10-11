package v1

import (
	"github.com/JeasonZuo/gochat/service/user_service"
	"github.com/JeasonZuo/gochat/utils/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterUserForm struct {
	Name            string `json:"name" binding:"required"`
	AvatarUrl       string `json:"avatar_url" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// @BasePath /api/v1

// @Summary 注册新用户
// Tags 用户相关
// @Accept json
// @Produce json
// @Param name body string true "用户名"
// @Param avatar_url body string true "用户头像"
// @Param password body string true "密码"
// @Param confirm_password body string true "确认密码"
// @success 200 {object} app.Response
// @Router /user/sign_up [post]
func RegisterUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form RegisterUserForm
	)

	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, 10001, err.Error(), nil)
		return
	}

	userService := user_service.User{
		Name:      form.Name,
		AvatarUrl: form.AvatarUrl,
		Password:  form.Password,
	}

	id, err := userService.RegisterUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 10002, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, 10000, "ok", gin.H{
		"tt_number": id + 1000000,
	})
}
