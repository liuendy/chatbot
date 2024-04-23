package web

import (
	"chatbot/internal/application/service"
	"chatbot/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ChatController 控制聊天相关的 HTTP 请求
type ChatController struct {
	chatService *service.ChatService
}

// NewChatController 创建新的聊天控制器
func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{
		chatService: chatService,
	}
}

// StartChatSession 启动新的聊天会话
func (cc *ChatController) StartChatSession(c *gin.Context) {
	userID := c.Param("userID")
	session, err := cc.chatService.StartNewSession(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetChatSession 获取指定的聊天会话详情
func (cc *ChatController) GetChatSession(c *gin.Context) {
	sessionID := c.Param("sessionID")
	session, err := cc.chatService.GetSessionByID(c, sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	c.JSON(http.StatusOK, session)
}

// SendMessage 发送消息
func (cc *ChatController) SendMessage(c *gin.Context) {
	var message model.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.chatService.SendMessage(c, message.SessionID, &message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Message sent"})
}

// GetMessageHistory 获取消息历史
func (cc *ChatController) GetMessageHistory(c *gin.Context) {
	sessionID := c.Param("sessionID")
	messages, err := cc.chatService.GetMessageHistory(c, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// GetMessageHistory 获取消息历史
func (cc *ChatController) DeleteMessage(c *gin.Context) {
	sessionId := c.Query("sessionId")
	messageId := c.Query("messageId")
	err := cc.chatService.DeleteMessage(c, sessionId, messageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}
