package application

import (
	domainC "tienda/src/Client/Domain"
	"tienda/src/Notification/Application/services"
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type PostNotificationUseCase struct {
	notificationRepo domain.NotificationRepository
	clientRepo       domainC.ClientRepository
	rabitService     services.RabbitMQService
}

func NewPostNotificationUseCase(notificationRepo domain.NotificationRepository, clientRepo domainC.ClientRepository,
	rabitSerivce services.RabbitMQService) *PostNotificationUseCase {
	return &PostNotificationUseCase{notificationRepo: notificationRepo, clientRepo: clientRepo,
		rabitService: rabitSerivce}
}

func (uc *PostNotificationUseCase) Execute(notification entities.Notification) error {
	client, err := uc.clientRepo.Search(notification.ClientID)
	if err != nil {
		return err
	}
	uc.rabitService.SendMessage(notification.Content, notification.ClientID)
	return uc.notificationRepo.Send(notification, client)
}

//inyectar el seevicio de rabbitmq
