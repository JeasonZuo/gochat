package user_service

import (
	"github.com/JeasonZuo/gochat/models"
	"github.com/JeasonZuo/gochat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint
	Name      string
	AvatarUrl string
	Password  string
}

// 用户注册
func (u *User) RegisterUser() (uint, error) {
	user := map[string]any{
		"id":         u.ID,
		"name":       u.Name,
		"avatar_url": u.AvatarUrl,
		"password":   u.Password,
	}
	id, err := models.AddUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 用户登录
func (u *User) LoginUser() (string, error) {
	//获取用户信息
	user, err := models.GetUserById(u.ID)
	if err != nil {
		return "", err
	}

	//校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return "", err
	}

	//生成token
	token, err := utils.GenerateToken(user.ID, user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}
