package routes

//mapper
import (
	"tienda/src/products/infraestructure/controller"

	gin "github.com/gin-gonic/gin"
)

// mejorr rutas
func SetRoutes(router *gin.Engine, postController *controller.PostProductsHandler,
	getController *controller.GetProductsHandler, getOneController *controller.GetOneProductHandler, deleteProduct *controller.DeleteProductHandler,
	putController *controller.PutProductHandler, shortgHandle *controller.GetNewProductsHanlderShort, longHandle *controller.LongPollingHandler) {
	router.GET("/products", getController.Handle)
	router.POST("/products", postController.Handle)
	router.GET("/products/:id", getOneController.HandleGetOne)
	router.DELETE("/products/:id", deleteProduct.HandleDelete)
	router.PUT("/products/:id", putController.HandlePut)
	router.GET("/products/new", shortgHandle.HandleLong)    // Short polling endpoint
	router.GET("/products/longpoll", longHandle.HandleLong) // Long polling endpoint

}
