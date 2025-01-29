package controller

import (
	"net/http"
	"strconv"
	application "tienda/src/Notification/Application"

	"github.com/gin-gonic/gin"
)

type DeleteNotificationHandler struct {
	DeleteNotificationUseCase *application.DeleteNotificationUseCase
}

func NewDeleteNotificationHandler(deleteNotificationUseCase *application.DeleteNotificationUseCase) *DeleteNotificationHandler {
	return &DeleteNotificationHandler{DeleteNotificationUseCase: deleteNotificationUseCase}
}

func (h *DeleteNotificationHandler) HandleDelete(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
		return
	}

	err = h.DeleteNotificationUseCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Mensaje eliminada"})
}
