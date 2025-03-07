package controller

import (
	"net/http"
	applicationnegocio "tienda/src/products/application-negocio"

	"github.com/gin-gonic/gin"
)

type GetProductsHandler struct {
	GetAllProductsUseCase *applicationnegocio.GetAllProductsUseCase
}

func NewGetProductsHandler(getAllProductsUseCase *applicationnegocio.GetAllProductsUseCase) *GetProductsHandler {
	return &GetProductsHandler{GetAllProductsUseCase: getAllProductsUseCase}
}

func (handle *GetProductsHandler) Handle(g *gin.Context) {
	products, err := handle.GetAllProductsUseCase.Execute()

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"Lista de productos": products})
}
