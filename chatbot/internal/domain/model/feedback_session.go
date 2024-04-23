package model

// FeedbackSession 聚合根，代表一系列用户反馈
type FeedbackSession struct {
	SessionID     string          `bson:"sessionId"`     // 反馈会话的唯一标识
	UserID        string          `bson:"userId"`        // 用户标识
	FeedbackItems []*FeedbackItem `bson:"feedbackItems"` // 反馈项列表
}

// NewFeedbackSession 创建新的反馈会话
func NewFeedbackSession(sessionID, userID string) *FeedbackSession {
	return &FeedbackSession{
		SessionID:     sessionID,
		UserID:        userID,
		FeedbackItems: []*FeedbackItem{},
	}
}

// AddFeedbackItem 向会话中添加一个新的反馈项
func (fs *FeedbackSession) AddFeedbackItem(item *FeedbackItem) {
	fs.FeedbackItems = append(fs.FeedbackItems, item)
}
