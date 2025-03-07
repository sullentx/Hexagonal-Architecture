package main

import (
	"log"
	infraClient "tienda/src/client/infraestructure"
	routesClient "tienda/src/client/infraestructure/routes"
	infra "tienda/src/notification/infraestructure"
	"tienda/src/notification/infraestructure/adapters"
	routesNotification "tienda/src/notification/infraestructure/routes"
	infraestructure "tienda/src/products/infraestructure"
	routes "tienda/src/products/infraestructure/routes"

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

	infraClient.Init()
	infra.Init()
	routesClient.SetRoutes(router, infraClient.CclientSearcHandler, infraClient.CreateClientHandler, infraClient.DeleteClientHandler, infraClient.GetClientsHandler,
		infraClient.ModifyClientHandler)
	//inicializar rutas
	routesNotification.SetRoutes(router, infra.OPostNotificationHandler,
		infra.ODeleteNotificationHadler, infra.OSearchNotificationHadler,
		infra.OModifyNotificationHadler, infra.OGetNotificationHandler, infra.OShortPutNotificationHadler)

	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler, infraestructure.GetNewProductsHandler, infraestructure.LongPollingHandler)
	// conexion rabbit
	rabbitMQAdapter := adapters.InitRabbitMQ()
	defer rabbitMQAdapter.Close()

	log.Println("Server started at :8080")
	log.Fatal(router.Run())
}
