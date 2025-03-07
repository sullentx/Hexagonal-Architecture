package controller

import (
	"net/http"
	"strconv"
	application "tienda/src/client/aplication"
	entities "tienda/src/client/domain/entities"

	"github.com/gin-gonic/gin"
)

type ModifyClientHandler struct {
	ModifyClientUseCase *application.ModifyClientUseCase
}

func NewModifyClientHandler(modifyNotificationUseCase *application.ModifyClientUseCase) *ModifyClientHandler {
	return &ModifyClientHandler{ModifyClientUseCase: modifyNotificationUseCase}
}

func (h *ModifyClientHandler) HandleModify(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
		return
	}

	var client entities.Client
	err = g.BindJSON(&client)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.ModifyClientUseCase.Execute(id, client)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Cliente modificado"})
}
