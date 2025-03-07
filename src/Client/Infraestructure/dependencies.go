package infraestructure

import (
	application "tienda/src/client/aplication"
	controller "tienda/src/client/infraestructure/controller"
	core "tienda/src/core"
)

var (
	CclientSearcHandler *controller.SearchClientHandler
	CreateClientHandler *controller.CreateClientHandler
	DeleteClientHandler *controller.DeleteClientHandler
	GetClientsHandler   *controller.GetClientsHandler
	ModifyClientHandler *controller.ModifyClientHandler
)

func Init() {
	core.InitPostgres()
	db := core.GetDB()

	clientRepo := NewPostgresClientRepository(db)

	createClientUseCase := application.NewCreateClientUseCase(clientRepo)
	CreateClientHandler = controller.NewCreateClientHandler(createClientUseCase)

	deleteClientUseCase := application.NewDeleteClient(clientRepo)
	DeleteClientHandler = controller.NewDeleteClientHandler(deleteClientUseCase)

	getByIdClientUseCase := application.NewGetByIdClientUseCase(clientRepo)
	CclientSearcHandler = controller.NewSearchClient(getByIdClientUseCase)

	getClientsUseCase := application.NewGetClients(clientRepo)
	GetClientsHandler = controller.NewGetClientHandler(getClientsUseCase)

	modifyClientUseCase := application.NewModifyClient(clientRepo)
	ModifyClientHandler = controller.NewModifyClientHandler(modifyClientUseCase)
}
