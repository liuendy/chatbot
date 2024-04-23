package repository

import (
	"chatbot/internal/domain/model"
	"context"
)

// ChatSessionRepository 提供了对聊天会话的存储和检索操作的接口定义
type ChatSessionRepository interface {
	Save(ctx context.Context, session *model.ChatSession) error
	FindByID(ctx context.Context, sessionID string) (*model.ChatSession, error)
	AddMessage(ctx context.Context, sessionID string, message model.Message) error
	GetMessageHistory(ctx context.Context, sessionID string) ([]model.Message, error)
	RemoveMessage(ctx context.Context, sessionID string, messageID string) error
}
