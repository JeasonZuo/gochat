package v1

import (
	"github.com/JeasonZuo/gochat/pkg/app"
	"github.com/JeasonZuo/gochat/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var startTTNumber uint = 100000

type RegisterUserForm struct {
	Name            string `json:"name" binding:"required,min=2,max=10"`
	AvatarUrl       string `json:"avatar_url" binding:"required,max=255"`
	Password        string `json:"password" binding:"required,min=6,max=72"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// @BasePath /api/v1

// UserRegister
// @Tags 用户相关
// @Summary 注册新用户
// @Accept json
// @Produce json
// @Param name body string true "用户名"
// @Param avatar_url body string true "用户头像"
// @Param password body string true "密码"
// @Param confirm_password body string true "确认密码"
// @success 200 {object} app.Response
// @Router /sign_up [post]
func UserRegister(c *gin.Context) {
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
		"tt_number": id + startTTNumber,
	})
}

type LoginForm struct {
	TTNumber uint   `json:"tt_number" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=72"`
}

// UserLogin
// @Tags 用户相关
// @Summary 用户登陆
// @Accept json
// @Produce json
// @Param tt_number body int true "TT号"
// @Param password body string true "密码"
// @success 200 {object} app.Response
// @Router /sign_in [post]
func UserLogin(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form LoginForm
	)

	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, 10001, err.Error(), nil)
		return
	}

	userService := user_service.User{
		ID:       form.TTNumber - startTTNumber,
		Password: form.Password,
	}

	token, err := userService.LoginUser()
	if err != nil {
		appG.Response(http.StatusOK, 10002, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, 10000, "ok", gin.H{
		"token": token,
	})
}

type AddFriendForm struct {
	TTNumber uint `json:"tt_number" binding:"required"`
}

// AddFriend
// @Tags 用户相关
// @Summary 添加好友
// @Accept json
// @Produce json
// @Param tt_number body int true "TT号"
// @success 200 {object} app.Response
// @Router /add_friend [post]
func AddFriend(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddFriendForm
	)

	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, 10001, err.Error(), nil)
		return
	}

	loginUserId := c.GetUint("loginUserId")
	if loginUserId == 0 {
		appG.Response(http.StatusOK, 10002, "请登陆", nil)
		return
	}

	userService := user_service.User{
		ID: loginUserId,
	}
	err := userService.AddFriend(form.TTNumber - startTTNumber)
	if err != nil {
		appG.Response(http.StatusOK, 10002, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, 10000, "ok", nil)
}

func GetFriendList(c *gin.Context) {
	var appG = app.Gin{C: c}
	loginUserId := c.GetUint("loginUserId")

	userService := user_service.User{
		ID: loginUserId,
	}
	list, err := userService.GetFriendsList()
	if err != nil {
		appG.Response(http.StatusOK, 10002, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, 10000, "ok", list)
}
