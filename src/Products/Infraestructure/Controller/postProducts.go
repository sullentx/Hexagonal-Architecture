package controller

import (
	"net/http"
	applicationnegocio "tienda/src/products/application-negocio"
	entities "tienda/src/products/domain-negocio/entities"

	"github.com/gin-gonic/gin"
)

type PostProductsHandler struct {
	createProductUseCase *applicationnegocio.CreateProductUseCase
}

// constructor que devuelve un puntero a la estructura PostProductsHandler
func NewPostProductsHandler(createProductUseCase *applicationnegocio.CreateProductUseCase) *PostProductsHandler {
	return &PostProductsHandler{createProductUseCase: createProductUseCase}
}

func (handle *PostProductsHandler) Handle(g *gin.Context) {
	var product entities.Product //se crea una variable product de tipo Product

	if err := g.ShouldBind(&product); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	//se ejecuta el metodo Execute de la estructura CreateProductUseCase
	if err := handle.createProductUseCase.Execute(product); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, gin.H{"Message": "Product Creado"})
}

/*
func (handle *PostProductsHandler) LongPollNewProducts(g *gin.Context) {
	lastProductID, err := strconv.Atoi(g.Query("lastProductID"))
	if err != nil {
		lastProductID = 0
	}

	for {
		products, latestID, err := handle.getNewProductsUseCase.Execute(lastProductID)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		if len(products) > 0 {
			g.JSON(http.StatusOK, gin.H{"products": products, "lastProductID": latestID})
			return
		}

		time.Sleep(2 * time.Second) // Espera antes de volver a checar
	}
}

*/
