package repository

import (
	"chatbot/internal/domain/model"
	"context"
)

// FeedbackSessionRepository 提供了对反馈会话的存储和检索操作的接口定义
type FeedbackSessionRepository interface {
	Save(ctx context.Context, session *model.FeedbackSession) error
	FindByID(ctx context.Context, sessionID string) (*model.FeedbackSession, error)
}
