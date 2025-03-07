package controller

import (
	"net/http"
	application "tienda/src/client/aplication"

	"github.com/gin-gonic/gin"
)

type GetClientsHandler struct {
	GetClientsUseCase *application.GetClientsUseCase
}

func NewGetClientHandler(usecase *application.GetClientsUseCase) *GetClientsHandler {
	return &GetClientsHandler{GetClientsUseCase: usecase}
}

func (h *GetClientsHandler) HandleGetAll(g *gin.Context) {
	clients, err := h.GetClientsUseCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"Clientes": clients})
}
