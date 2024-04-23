package service

import (
	"chatbot/internal/domain/model"
	"chatbot/internal/domain/repository"
	"context"
)

// ChatService 应用服务
type ChatService struct {
	repo repository.ChatSessionRepository
}

// NewChatService 创建新的聊天服务
func NewChatService(repo repository.ChatSessionRepository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

// StartNewSession 启动新的聊天会话
func (s *ChatService) StartNewSession(ctx context.Context, userID string) (*model.ChatSession, error) {
	session := model.NewChatSession(userID)
	err := s.repo.Save(ctx, session)
	return session, err
}

// GetSessionByID 根据 ID 获取聊天会话
func (s *ChatService) GetSessionByID(ctx context.Context, id string) (*model.ChatSession, error) {
	session, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *ChatService) SendMessage(ctx context.Context, sessionID string, message *model.Message) error {
	return s.repo.AddMessage(ctx, sessionID, *message)
}

func (s *ChatService) GetMessageHistory(ctx context.Context, sessionID string) ([]model.Message, error) {
	return s.repo.GetMessageHistory(ctx, sessionID)
}

func (cs *ChatService) DeleteMessage(ctx context.Context, sessionID, messageID string) error {
	return cs.repo.RemoveMessage(ctx, sessionID, messageID)
}
