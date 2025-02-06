package controller

import (
	"net/http"
	"strconv"
	application "tienda/src/Notification/Application"
	entities "tienda/src/Notification/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type ModifyNotificationHandler struct {
	ModifyNotificationUseCase *application.ModifyNotificationUseCase
}

func NewModifyNotificationHandler(modifyNotificationUseCase *application.ModifyNotificationUseCase) *ModifyNotificationHandler {
	return &ModifyNotificationHandler{ModifyNotificationUseCase: modifyNotificationUseCase}
}

func (h *ModifyNotificationHandler) HandleModify(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	var notification entities.Notification
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": "debe de ser numero entero"})
		return
	}

	if err := g.ShouldBindJSON(&notification); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": "datos invalidos"})
		return
	}
	err = h.ModifyNotificationUseCase.Execute(id, notification.Content)
	if err == nil {
		g.JSON(http.StatusOK, gin.H{"message": "Notificación actualizada"})
	}
}
