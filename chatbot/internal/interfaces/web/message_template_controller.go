package web

import (
	"chatbot/internal/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageTemplateController struct {
	templateService *service.MessageTemplateService
}

func NewMessageTemplateController(templateService *service.MessageTemplateService) *MessageTemplateController {
	return &MessageTemplateController{
		templateService: templateService,
	}
}

func (mtc *MessageTemplateController) CreateTemplate(c *gin.Context) {
	var request struct {
		TemplateId   string `json:"templateId"`
		Content      string `json:"content"`
		UsageContext string `json:"usageContext"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	template, err := mtc.templateService.CreateTemplate(c, request.TemplateId, request.Content, request.UsageContext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, template)
}

// GetTemplate 根据模板ID获取模板详情
func (mtc *MessageTemplateController) GetTemplate(c *gin.Context) {
	templateID := c.Param("templateID")
	template, err := mtc.templateService.GetTemplate(c, templateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}
	c.JSON(http.StatusOK, template)
}

// DeleteTemplate 删除指定的模板
func (mtc *MessageTemplateController) DeleteTemplate(c *gin.Context) {
	templateID := c.Param("templateID")
	if err := mtc.templateService.DeleteTemplate(c, templateID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetAllTemplates 获取所有模板列表
func (mtc *MessageTemplateController) GetAllTemplates(c *gin.Context) {
	templates, err := mtc.templateService.GetAllTemplates(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, templates)
}
