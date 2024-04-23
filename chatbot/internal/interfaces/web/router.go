package web

import (
	"chatbot/internal/application/service"
	"github.com/gin-gonic/gin"
)

// NewRouter 创建并配置新的 Gin 路由器
func NewRouter(chatService *service.ChatService, feedbackService *service.FeedbackService, templateService *service.MessageTemplateService) *gin.Engine {
	router := gin.Default()

	// 聊天相关路由
	chatController := NewChatController(chatService)
	router.POST("/start_chat/:userID", chatController.StartChatSession)
	router.GET("/get_chat/:sessionID", chatController.GetChatSession)

	// 消息模板相关路由
	templateController := NewMessageTemplateController(templateService)
	router.POST("/templates", templateController.CreateTemplate)
	router.GET("/templates/:templateID", templateController.GetTemplate)
	router.DELETE("/templates/:templateID", templateController.DeleteTemplate)
	router.GET("/templates", templateController.GetAllTemplates)

	// 反馈相关路由
	feedbackController := NewFeedbackController(feedbackService)
	router.POST("/feedback", feedbackController.AddFeedback)
	router.GET("/feedback/:sessionID", feedbackController.GetFeedbackSession)

	// 消息相关路由
	router.POST("/messages", chatController.SendMessage)
	router.GET("/messages/:sessionID", chatController.GetMessageHistory)
	router.GET("/delMessage", chatController.DeleteMessage)

	return router
}
