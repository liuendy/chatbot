package service

import (
	"chatbot/internal/domain/model"
	"chatbot/internal/domain/repository"
	"context"
)

type FeedbackService struct {
	repo repository.FeedbackSessionRepository
}

// NewFeedbackService 创建一个新的 FeedbackService 实例
func NewFeedbackService(repo repository.FeedbackSessionRepository) *FeedbackService {
	return &FeedbackService{repo: repo}
}

// AddFeedback 将用户反馈添加到指定的反馈会话
func (f *FeedbackService) AddFeedback(ctx context.Context, sessionID, userID string, content string, rating int) (*model.FeedbackSession, error) {
	session, err := f.repo.FindByID(ctx, sessionID)
	if err != nil {
		session = &model.FeedbackSession{
			SessionID: sessionID,
			UserID:    userID,
		}
	}

	// 创建一个新的反馈项
	feedbackItem := model.NewFeedbackItem(model.GenerateID(), content, rating)
	session.AddFeedbackItem(feedbackItem)

	// 保存更新后的会话
	return session, f.repo.Save(ctx, session)
}

// GetFeedbackSession 根据会话ID获取反馈会话
func (f *FeedbackService) GetFeedbackSession(ctx context.Context, sessionID string) (*model.FeedbackSession, error) {
	return f.repo.FindByID(ctx, sessionID)
}
