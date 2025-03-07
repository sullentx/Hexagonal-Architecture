package routes

import (
	"tienda/src/client/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, searchClient *controller.SearchClientHandler, createClient *controller.CreateClientHandler,
	deleteClient *controller.DeleteClientHandler, getClients *controller.GetClientsHandler, modifyClient *controller.ModifyClientHandler) {
	router.GET("/clients/:id", searchClient.HandleSearch)
	router.POST("/clients", createClient.HandleCreate)
	router.DELETE("/clients/:id", deleteClient.HandleDelete)
	router.GET("/clients", getClients.HandleGetAll)
	router.PUT("/clients/:id", modifyClient.HandleModify)
}
