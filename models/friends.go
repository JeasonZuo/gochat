package models

import (
	"fmt"
	"gorm.io/gorm"
)

type FriendModel struct {
	gorm.Model
	UserId       uint      `json:"user_id" gorm:"index"`
	FriendUserId uint      `json:"friend_user_id"`
	FriendUser   UserModel `gorm:"foreignKey:ID;references:friend_user_id"`
}

type FriendInfo struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func (f *FriendModel) TableName() string {
	return "tb_friends"
}

// GetFriendsList 获取好友列表
func GetFriendsList(userId uint) []*FriendInfo {
	var friends []*FriendInfo
	db.Model(&FriendModel{}).Select("tb_users.id, tb_users.name, tb_users.avatar_url").
		Joins("inner join tb_users on tb_friends.friend_user_id = tb_users.id").
		Where("tb_friends.user_id = ?", userId).
		Scan(&friends)

	return friends
}

// AddFriend 添加好友
func AddFriend(userId, friendUserId uint) error {
	friendModel := FriendModel{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
	friend := FriendModel{}
	r := db.Model(&FriendModel{}).Where(friendModel).Limit(1).Find(&friend)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected > 0 {
		return fmt.Errorf("用户已经是好友")
	}

	return db.Create(&FriendModel{UserId: userId, FriendUserId: friendUserId}).Error
}
