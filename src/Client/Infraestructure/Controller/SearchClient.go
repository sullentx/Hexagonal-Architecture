package controller

import (
	"net/http"
	"strconv"
	aplication "tienda/src/client/aplication"

	"github.com/gin-gonic/gin"
)

type SearchClientHandler struct {
	SearchClientUseCase *aplication.GetByIdClientUseCase
}

func NewSearchClient(uc *aplication.GetByIdClientUseCase) *SearchClientHandler {
	return &SearchClientHandler{SearchClientUseCase: uc}
}

func (h *SearchClientHandler) HandleSearch(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
	}

	client, err := h.SearchClientUseCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"Cliente": client})
}
