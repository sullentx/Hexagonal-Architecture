package controller

import (
	"net/http"
	"strconv"
	application "tienda/src/Notification/Application"

	"github.com/gin-gonic/gin"
)

type SearchNotificationHandler struct {
	SearchNotificationUseCase *application.SearchNotificationUseCase
}

func NewSearchNotificationHandler(searchNotificationUseCase *application.SearchNotificationUseCase) *SearchNotificationHandler {
	return &SearchNotificationHandler{SearchNotificationUseCase: searchNotificationUseCase}
}

func (h *SearchNotificationHandler) HandleSearch(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
	}

	notification, err := h.SearchNotificationUseCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"Notificacion": notification})
}
