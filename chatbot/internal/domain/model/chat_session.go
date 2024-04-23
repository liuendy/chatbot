package model

import "time"

// ChatSession 聊天会话聚合根
type ChatSession struct {
	ID       string    `bson:"id"`
	UserID   string    `bson:"userId"`
	Messages []Message `bson:"messages"`
}

// Message 实体
type Message struct {
	ID        string    `bson:"id"`
	SessionID string    `bson:"sessionId"`
	Content   string    `bson:"content"`
	Sender    string    `bson:"sender"`
	TimeSent  time.Time `bson:"timeSent"`
}

// NewChatSession 创建新的聊天会话
func NewChatSession(userID string) *ChatSession {
	return &ChatSession{
		ID:       GenerateID(),
		UserID:   userID,
		Messages: []Message{},
	}
}

// AddMessage 向聊天会话中添加消息
func (cs *ChatSession) AddMessage(content string) {
	message := Message{
		ID:       GenerateID(),
		Content:  content,
		TimeSent: time.Now(),
	}
	cs.Messages = append(cs.Messages, message)
}
