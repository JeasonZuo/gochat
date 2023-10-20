package models

import (
	"gorm.io/gorm"
	"time"
)

type MessageModel struct {
	gorm.Model
	FromUserId uint   `json:"from_user_id" gorm:"index"`
	ToUserId   uint   `json:"to_user_id" gorm:"index"`
	Content    string `json:"content" gorm:"size:2300"`
}

type MessageInfo struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	FromUserId uint      `json:"from_user_id"`
	ToUserId   uint      `json:"to_user_id"`
	Content    string    `json:"content"`
}

func (m *MessageModel) TableName() string {
	return "tb_messages"
}

// InsertMessage 写入信息
func InsertMessage(data map[string]any) error {
	messageModel := MessageModel{
		FromUserId: data["from_user_id"].(uint),
		ToUserId:   data["to_user_id"].(uint),
		Content:    data["content"].(string),
	}
	return db.Create(&messageModel).Error
}

// GetMessageList 读取信息列表
func GetMessageList(fromUserId, toUserId uint) ([]*MessageInfo, error) {
	messages := make([]*MessageInfo, 0)
	err := db.Model(&MessageModel{}).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Or("to_user_id = ? AND  from_user_id= ?", fromUserId, toUserId).Find(&messages).Error
	return messages, err
}
