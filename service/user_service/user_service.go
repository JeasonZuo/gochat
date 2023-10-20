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

// RegisterUser 用户注册
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

// LoginUser 用户登录
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

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (*models.UserInfo, error) {
	userModel, err := models.GetUserById(u.ID)
	if err != nil {
		return nil, err
	}

	userInfo := &models.UserInfo{
		ID:        userModel.ID,
		Name:      userModel.Name,
		AvatarUrl: userModel.AvatarUrl,
	}
	return userInfo, nil
}

// AddFriend 添加好友
func (u *User) AddFriend(friendId uint) error {
	err := models.AddFriend(u.ID, friendId)
	if err != nil {
		return err
	}

	//自动互相添加为好友
	return models.AddFriend(friendId, u.ID)
}

// 获取好友列表
func (u *User) GetFriendsList() ([]*models.UserInfo, error) {
	list := models.GetFriendsList(u.ID)
	return list, nil
}
