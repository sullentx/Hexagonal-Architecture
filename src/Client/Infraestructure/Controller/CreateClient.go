package controller

import (
	application "tienda/src/Client/Aplication"
	entities "tienda/src/Client/Domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateClientHandler struct {
	CreateClientUseCase *application.CreateClientUseCase
}

func NewCreateClientHandler(usecase *application.CreateClientUseCase) *CreateClientHandler {
	return &CreateClientHandler{CreateClientUseCase: usecase}
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
