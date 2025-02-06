package controller

import (
	"net/http"
	"strconv"
	applicationnegocio "tienda/src/Products/Application-negocio"

	"github.com/gin-gonic/gin"
)

type GetNewProductsHanlderShort struct {
	GetNewProductsUseCase *applicationnegocio.GetNewProductsUseCase
}

func NewGetNewProductsHandler(getNewProductsUseCase *applicationnegocio.GetNewProductsUseCase) *GetNewProductsHanlderShort {
	return &GetNewProductsHanlderShort{GetNewProductsUseCase: getNewProductsUseCase}
}

func (h *GetNewProductsHanlderShort) HandleLong(c *gin.Context) {
	lastProductID, err := strconv.Atoi(c.Query("lastProductID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lastProductID"})
		return
	}

	products, latestID, err := h.GetNewProductsUseCase.Execute(lastProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products, "latestID": latestID})
}
