package infraestructure

import (
	infraestructure "tienda/src/client/infraestructure"
	core "tienda/src/core"
	application "tienda/src/notification/application"
	"tienda/src/notification/application/repositories"
	"tienda/src/notification/application/services"
	"tienda/src/notification/infraestructure/adapters"
	controller "tienda/src/notification/infraestructure/controller"
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
