package web

import (
	"chatbot/internal/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FeedbackController 控制用户反馈相关的 HTTP 请求
type FeedbackController struct {
	feedbackService *service.FeedbackService
}

// NewFeedbackController 创建新的反馈控制器
func NewFeedbackController(feedbackService *service.FeedbackService) *FeedbackController {
	return &FeedbackController{
		feedbackService: feedbackService,
	}
}

// AddFeedback 接收和处理用户反馈
func (fc *FeedbackController) AddFeedback(c *gin.Context) {
	var request struct {
		SessionID string `json:"sessionId"`
		UserID    string `json:"userId"`
		Content   string `json:"content"`
		Rating    int    `json:"rating"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := fc.feedbackService.AddFeedback(c, request.SessionID, request.UserID, request.Content, request.Rating);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetFeedbackSession 获取指定的反馈会话
func (fc *FeedbackController) GetFeedbackSession(c *gin.Context) {
	sessionID := c.Param("sessionID")
	session, err := fc.feedbackService.GetFeedbackSession(c, sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	c.JSON(http.StatusOK, session)
}
