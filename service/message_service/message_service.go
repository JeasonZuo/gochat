package message_service

import (
	"github.com/JeasonZuo/gochat/models"
)

type Message struct {
	ID         int64  `json:"id"`
	CreateTime string `json:"createTime"`
	Content    string `json:"content"`
	FromUserId uint   `json:"from_user_id"`
	ToUserId   uint   `json:"to_user_id"`
}

// SaveMessage 存储消息
func (m *Message) SaveMessage() error {
	message := map[string]any{
		"from_user_id": m.FromUserId,
		"to_user_id":   m.ToUserId,
		"content":      m.Content,
	}
	return models.InsertMessage(message)
}

// GetMessageList 获取消息列表
func (m *Message) GetMessageList() ([]*models.MessageInfo, error) {
	return models.GetMessageList(m.FromUserId, m.ToUserId)
}
