package user_service

import (
	"github.com/JeasonZuo/gochat/models"
)

type User struct {
	ID        int
	Name      string
	AvatarUrl string
	Password  string
}

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
