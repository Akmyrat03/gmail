package handler

import (
	"context"
	"mail-service/models"
	"mail-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	Service *service.ContactService
}

func NewContactHandler(service *service.ContactService) *ContactHandler {
	return &ContactHandler{Service: service}
}

func (h *ContactHandler) SendMessage(c *gin.Context) {
	var message models.ContactMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	ctx := context.Background()
	if err := h.Service.SendMessage(ctx, message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mesaj gönderilemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mesaj başarıyla gönderildi"})
}
