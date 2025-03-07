package controller

import (
	application "tienda/src/client/aplication"
	entities "tienda/src/client/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateClientHandler struct {
	CreateClientUseCase *application.CreateClientUseCase
}

func NewCreateClientHandler(createClientUseCase *application.CreateClientUseCase) *CreateClientHandler {
	return &CreateClientHandler{CreateClientUseCase: createClientUseCase}
}

func (h *CreateClientHandler) HandleCreate(g *gin.Context) {
	var client entities.Client
	if err := g.ShouldBindJSON(&client); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.CreateClientUseCase.Execute(client)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(201, gin.H{"message": "Cliente creado"})
}
