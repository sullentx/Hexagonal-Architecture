package routes

import (
	controller "tienda/src/Client/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, searchClient *controller.SearchClientHandler) {
	router.GET("/clients/:id", searchClient.HandleSearch)
}
