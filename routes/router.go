package routes

import (
	"mail-service/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, contactHandler *handler.ContactHandler) {
	api := r.Group("/api")
	{
		api.POST("/contact", contactHandler.SendMessage)
	}
}
