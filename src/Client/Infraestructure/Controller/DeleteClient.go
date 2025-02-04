package controller

import (
	"net/http"
	"strconv"
	application "tienda/src/Client/Aplication"

	"github.com/gin-gonic/gin"
)

type DeleteClientHandler struct {
	DeleteClientUseCase *application.DeleteClientUseCase
}

func NewDeleteClientHandler(uc *application.DeleteClientUseCase) *DeleteClientHandler {
	return &DeleteClientHandler{DeleteClientUseCase: uc}
}

func (h *DeleteClientHandler) HandleDelete(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
		return
	}

	err = h.DeleteClientUseCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado"})
}
