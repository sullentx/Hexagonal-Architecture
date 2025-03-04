package main

import (
	"log"
	infraClient "tienda/src/Client/Infraestructure"
	routesClient "tienda/src/Client/Infraestructure/routes"
	infra "tienda/src/Notification/Infraestructure"
	routesNotification "tienda/src/Notification/Infraestructure/routes"
	infraestructure "tienda/src/Products/Infraestructure"
	routes "tienda/src/Products/Infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Inicializar dependencias
	infraestructure.Init()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	router.Use(cors.New(config))
	infra.Init()
	infraClient.Init()
	routesClient.SetRoutes(router, infraClient.CclientSearcHandler, infraClient.CreateClientHandler, infraClient.DeleteClientHandler, infraClient.GetClientsHandler,
		infraClient.ModifyClientHandler)
	//inicializar rutas
	routesNotification.SetRoutes(router, infra.OPostNotificationHandler,
		infra.ODeleteNotificationHadler, infra.OSearchNotificationHadler,
		infra.OModifyNotificationHadler, infra.OGetNotificationHandler, infra.OShortPutNotificationHadler)

	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler, infraestructure.GetNewProductsHandler, infraestructure.LongPollingHandler)
	// Iniciar el servidor

	log.Println("Server started at :8080")
	log.Fatal(router.Run())
}
