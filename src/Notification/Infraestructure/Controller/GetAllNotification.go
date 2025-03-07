package controller

import (
	"net/http"
	application "tienda/src/notification/application"

	"github.com/gin-gonic/gin"
)

type GetAllNotificationHandler struct {
	GetAllNotificationUseCase *application.GetAllNotificationUseCase
}

func NewGetAllNotificationHandler(getAllNotificationUseCase *application.GetAllNotificationUseCase) *GetAllNotificationHandler {
	return &GetAllNotificationHandler{GetAllNotificationUseCase: getAllNotificationUseCase}

}

func (h *GetAllNotificationHandler) HandleGetAll(g *gin.Context) {
	notification, err := h.GetAllNotificationUseCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"Notificaciones": notification})
}
