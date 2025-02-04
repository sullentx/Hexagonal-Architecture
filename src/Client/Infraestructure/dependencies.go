package infraestructure

import (
	application "tienda/src/Client/Aplication"
	controller "tienda/src/Client/Infraestructure/Controller"
	core "tienda/src/Core"
)

var (
	CclientSearcHandler *controller.SearchClientHandler
)

func Init() {
	core.InitPostgres()
	db := core.GetDB()

	clientRepo := NewPostgresClientRepository(db)

	getByIdClientUseCase := application.NewGetByIdClientUseCase(clientRepo)
	CclientSearcHandler = controller.NewSearchClient(getByIdClientUseCase)
}
