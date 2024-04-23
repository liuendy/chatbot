package service

import (
	"chatbot/internal/domain/model"
	"chatbot/internal/domain/repository"
	"context"
)

type MessageTemplateService struct {
	repo repository.MessageTemplateRepository
}

// NewMessageTemplateService 创建一个新的消息模板服务
func NewMessageTemplateService(repo repository.MessageTemplateRepository) *MessageTemplateService {
	return &MessageTemplateService{
		repo: repo,
	}
}

// CreateTemplate 创建新的消息模板
func (mts *MessageTemplateService) CreateTemplate(ctx context.Context, templateId, content, usageContext string) (*model.MessageTemplate, error) {
	template := model.NewMessageTemplate(templateId, content, usageContext)
	if err := mts.repo.Save(ctx, template); err != nil {
		return nil, err
	}
	return template, nil
}

// GetTemplate 根据模板ID获取模板
func (mts *MessageTemplateService) GetTemplate(ctx context.Context, templateID string) (*model.MessageTemplate, error) {
	return mts.repo.FindByID(ctx, templateID)
}

// DeleteTemplate 删除模板
func (mts *MessageTemplateService) DeleteTemplate(ctx context.Context, templateID string) error {
	return mts.repo.Delete(ctx, templateID)
}

// GetAllTemplates 获取所有模板
func (mts *MessageTemplateService) GetAllTemplates(ctx context.Context) ([]*model.MessageTemplate, error) {
	return mts.repo.FindAll(ctx)
}
