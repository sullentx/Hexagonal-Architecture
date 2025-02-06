package infraestructure

import (
	core "tienda/src/Core"
	applicationnegocio "tienda/src/Products/Application-negocio"
	controller "tienda/src/Products/Infraestructure/Controller"
)

var (
	PostProductsHandler   *controller.PostProductsHandler
	GetProductsHandler    *controller.GetProductsHandler
	GetOneProductHadler   *controller.GetOneProductHandler
	DeleteProductHadler   *controller.DeleteProductHandler
	PutProductHadler      *controller.PutProductHandler
	GetNewProductsHandler *controller.GetNewProductsHanlderShort
	LongPollingHandler    *controller.LongPollingHandler
)

func Init() {
	// Inicializar la conexión a la base de datos
	core.InitPostgres()
	db := core.GetDB()

	// Crear instancias del repositorio y casos de uso
	productRepo := NewPostgresProductRepository(db)
	createProductUseCase := applicationnegocio.NewCreateProduct(productRepo)
	getAllProductsUseCase := applicationnegocio.GetAllProducts(productRepo)
	getOneProductUseCase := applicationnegocio.GetOneProduct(productRepo)
	deleteProductUseCase := applicationnegocio.DeleteProduct(productRepo)
	putProductUseCase := applicationnegocio.PutProduct(productRepo)
	getNewProductsUseCase := applicationnegocio.NewGetNewProductsUseCase(productRepo)

	// Crear instancias de los controladores
	PostProductsHandler = controller.NewPostProductsHandler(createProductUseCase)
	GetProductsHandler = controller.NewGetProductsHandler(getAllProductsUseCase)
	GetOneProductHadler = controller.NewGetOneProductHandler(getOneProductUseCase)
	DeleteProductHadler = controller.NewDeleteProductHandler(deleteProductUseCase)
	PutProductHadler = controller.NewPutProductUseCase(putProductUseCase)
	GetNewProductsHandler = controller.NewGetNewProductsHandler(getNewProductsUseCase)
	LongPollingHandler = controller.NewLongPollingHandler(getNewProductsUseCase)

}
