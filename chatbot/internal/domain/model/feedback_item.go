package model

import "time"

// FeedbackItem 表示用户的具体反馈信息
type FeedbackItem struct {
	ID        string    `bson:"id"`        // 反馈项的唯一标识
	Content   string    `bson:"content"`   // 反馈的具体内容
	Rating    int       `bson:"rating"`    // 反馈的评分
	CreatedAt time.Time `bson:"createdAt"` // 反馈的创建时间
}

// NewFeedbackItem 创建新的反馈项
func NewFeedbackItem(id, content string, rating int) *FeedbackItem {
	return &FeedbackItem{
		ID:        id,
		Content:   content,
		Rating:    rating,
		CreatedAt: time.Now(),
	}
}
