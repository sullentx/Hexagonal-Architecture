package infraestructure

import (
	infraestructure "tienda/src/Client/Infraestructure"
	core "tienda/src/Core"
	application "tienda/src/Notification/Application"
	controller "tienda/src/Notification/Infraestructure/Controller"
)

var (
	OPostNotificationHandler  *controller.PostNotificationHandler
	OGetNotificationHandler   *controller.GetAllNotificationHandler
	OSearchNotificationHadler *controller.SearchNotificationHandler
	ODeleteNotificationHadler *controller.DeleteNotificationHandler
	OModifyNotificationHadler *controller.ModifyNotificationHandler
)

func Init() {
	core.InitPostgres()
	db := core.GetDB()

	notificationRepo := NewPostgresNotificationRepository(db)
	clientRepo := infraestructure.NewPostgresClientRepository(db)

	SendNotificationUseCase := application.NewPostNotificationUseCase(notificationRepo, clientRepo)
	GetAllNotificationUseCase := application.GetAllNotification(notificationRepo)
	SearchNotificationUseCase := application.SearchNotification(notificationRepo)
	DeleteNotificationUseCase := application.DeleteNotification(notificationRepo)
	ModifyNotificationUseCase := application.ModifyNotification(notificationRepo)
	//Crear instancias de los controladores
	OPostNotificationHandler = controller.NewPostNotificationHandler(SendNotificationUseCase)
	OGetNotificationHandler = controller.NewGetAllNotificationHandler(GetAllNotificationUseCase)
	OSearchNotificationHadler = controller.NewSearchNotificationHandler(SearchNotificationUseCase)
	ODeleteNotificationHadler = controller.NewDeleteNotificationHandler(DeleteNotificationUseCase)
	OModifyNotificationHadler = controller.NewModifyNotificationHandler(ModifyNotificationUseCase)
}
