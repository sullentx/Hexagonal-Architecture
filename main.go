package main

import (
	"log"
	infra "tienda/src/Notification/Infraestructure"
	routesNotification "tienda/src/Notification/Infraestructure/routes"
	infraestructure "tienda/src/Products/Infraestructure"
	routes "tienda/src/Products/Infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Inicializar dependencias
	infraestructure.Init()
	infra.Init()
	//inicializar rutas
	routesNotification.SetRoutes(router, infra.OPostNotificationHandler,
		infra.ODeleteNotificationHadler, infra.OSearchNotificationHadler, infra.OModifyNotificationHadler, infra.OGetNotificationHandler)

	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler)
	// Iniciar el servidor

	log.Println("Server started at :8080")
	log.Fatal(router.Run())
}
