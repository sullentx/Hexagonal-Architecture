package infraestructure

import (
	infraestructure "tienda/src/Client/Infraestructure"
	core "tienda/src/Core"
	application "tienda/src/Notification/Application"
	"tienda/src/Notification/Application/repositories"
	"tienda/src/Notification/Application/services"
	controller "tienda/src/Notification/Infraestructure/Controller"
	"tienda/src/Notification/Infraestructure/adapters"
)

var (
	OPostNotificationHandler    *controller.PostNotificationHandler
	OGetNotificationHandler     *controller.GetAllNotificationHandler
	OSearchNotificationHadler   *controller.SearchNotificationHandler
	ODeleteNotificationHadler   *controller.DeleteNotificationHandler
	OModifyNotificationHadler   *controller.ModifyNotificationHandler
	OShortPutNotificationHadler *controller.GetNotificationsHandlerShort
)

func Init() {
	core.InitPostgres()
	db := core.GetDB()

	notificationRepo := NewPostgresNotificationRepository(db)
	clientRepo := infraestructure.NewPostgresClientRepository(db)
	rabbitMQAdapter := adapters.InitRabbitMQ()
	rabbitMQUseCase := repositories.NewRabbitMQUseCase(rabbitMQAdapter)
	rabbitMQService := services.NewRabbitMQService(rabbitMQUseCase)

	SendNotificationUseCase := application.NewPostNotificationUseCase(notificationRepo, clientRepo, *rabbitMQService)
	GetAllNotificationUseCase := application.GetAllNotification(notificationRepo)
	SearchNotificationUseCase := application.SearchNotification(notificationRepo)
	DeleteNotificationUseCase := application.DeleteNotification(notificationRepo)
	ModifyNotificationUseCase := application.ModifyNotification(notificationRepo)
	ShortPutNotificationUseCase := application.NewGetNotificationsUseCase(notificationRepo)
	//Crear instancias de los controladores
	OPostNotificationHandler = controller.NewPostNotificationHandler(SendNotificationUseCase)
	OGetNotificationHandler = controller.NewGetAllNotificationHandler(GetAllNotificationUseCase)
	OSearchNotificationHadler = controller.NewSearchNotificationHandler(SearchNotificationUseCase)
	ODeleteNotificationHadler = controller.NewDeleteNotificationHandler(DeleteNotificationUseCase)
	OModifyNotificationHadler = controller.NewModifyNotificationHandler(ModifyNotificationUseCase)
	OShortPutNotificationHadler = controller.NewGetNotificationsHandler(ShortPutNotificationUseCase)
}
