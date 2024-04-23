package repository

import (
	"chatbot/internal/domain/model"
	"context"
)

// MessageRepository 提供了对消息的存储和检索操作的接口定义
type MessageRepository interface {
	Save(ctx context.Context, message *model.Message) error
	FindBySessionID(ctx context.Context, sessionID string) ([]*model.Message, error)
	Delete(ctx context.Context, messageID string) error
}
