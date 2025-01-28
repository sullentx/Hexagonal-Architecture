package main

import (
	"log"
	infraestructure "tienda/src/Products/Infraestructure"
	routes "tienda/src/Products/Infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infraestructure.Init()

	router := gin.Default()
	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler)
	// Iniciar el servidor
	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
