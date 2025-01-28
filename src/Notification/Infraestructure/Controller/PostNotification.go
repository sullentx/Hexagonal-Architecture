package controller

import (
	"net/http"
	application "tienda/src/Notification/Application"
	entities "tienda/src/Notification/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type PostNotificationHandler struct {
	PostNotificationUseCase *application.PostNotificationUseCase
}

func NewPostNotificationHandler(postNotificationUseCase *application.PostNotificationUseCase) *PostNotificationHandler {
	return &PostNotificationHandler{PostNotificationUseCase: postNotificationUseCase}
}

func (h *PostNotificationHandler) HandlePost(g *gin.Context) {
	var notification entities.Notification

	if err := g.ShouldBind(&notification); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.PostNotificationUseCase.Execute(notification); err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, gin.H{"message": "Notificacion creada"})
}
