package controller

import (
	"net/http"
	"strconv"
	"time"

	application "tienda/src/Notification/Application"

	"github.com/gin-gonic/gin"
)

type GetNotificationsHandlerShort struct {
	GetNotificationsUseCase *application.GetNotificationsUseCase
}

func NewGetNotificationsHandler(useCase *application.GetNotificationsUseCase) *GetNotificationsHandlerShort {
	return &GetNotificationsHandlerShort{GetNotificationsUseCase: useCase}
}

func (h *GetNotificationsHandlerShort) HandleGet(c *gin.Context) {
	clientIDParam := c.Param("id")
	clientID, err := strconv.Atoi(clientIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id debe ser un número entero"})
		println(clientID)
		return
	}

	// Tiempo máximo de espera para short polling (5 segundos)
	pollingTimeout := time.Second * 5
	start := time.Now()

	for {
		notifications, err := h.GetNotificationsUseCase.Execute(clientID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(notifications) > 0 {
			c.JSON(http.StatusOK, notifications)
			return
		}

		// Si ya pasó el tiempo límite, responde con 204 No Content
		if time.Since(start) > pollingTimeout {
			c.Status(http.StatusNoContent)
			return
		}

		// Espera 500ms antes de volver a consultar
		time.Sleep(500 * time.Millisecond)
	}
}
