package repository

import (
	"chatbot/internal/domain/model"
	"context"
)

// MessageTemplateRepository 提供了对消息模板的存储和检索操作的接口定义
type MessageTemplateRepository interface {
	Save(ctx context.Context, template *model.MessageTemplate) error
	FindByID(ctx context.Context, templateID string) (*model.MessageTemplate, error)
	Delete(ctx context.Context, templateID string) error
	FindAll(ctx context.Context) ([]*model.MessageTemplate, error)
}
