package application

import (
	domainC "tienda/src/client/domain"
	"tienda/src/notification/application/services"
	domain "tienda/src/notification/domain"
	entities "tienda/src/notification/domain/entities"
)

type PostNotificationUseCase struct {
	notificationRepo domain.NotificationRepository
	clientRepo       domainC.ClientRepository
	rabitService     services.RabbitMQService
}

func NewPostNotificationUseCase(notificationRepo domain.NotificationRepository, clientRepo domainC.ClientRepository,
	rabitSerivce services.RabbitMQService) *PostNotificationUseCase {
	return &PostNotificationUseCase{
		notificationRepo: notificationRepo,
		clientRepo:       clientRepo,
		rabitService:     rabitSerivce}
}

func (uc *PostNotificationUseCase) Execute(notification entities.Notification) error {
	client, err := uc.clientRepo.Search(notification.ClientID)
	if err != nil {
		return err
	}
	uc.rabitService.SendMessage(notification.Content, notification.ClientID, client.Name)
	return nil
}

//inyectar el seevicio de rabbitmq
