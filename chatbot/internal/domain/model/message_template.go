package model

// MessageTemplate 表示一个消息模板
type MessageTemplate struct {
	TemplateID   string `bson:"templateId"`   // 模板的唯一标识符
	Content      string `bson:"content"`      // 模板内容
	UsageContext string `bson:"usageContext"` // 模板使用的上下文
}

// NewMessageTemplate 创建新的消息模板
func NewMessageTemplate(templateID, content, usageContext string) *MessageTemplate {
	return &MessageTemplate{
		TemplateID:   templateID,
		Content:      content,
		UsageContext: usageContext,
	}
}
