package application

import (
	domainC "tienda/src/Client/Domain"
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type PostNotificationUseCase struct {
	notificationRepo domain.NotificationRepository
	clientRepo       domainC.ClientRepository
}

func NewPostNotificationUseCase(notificationRepo domain.NotificationRepository, clientRepo domainC.ClientRepository) *PostNotificationUseCase {
	return &PostNotificationUseCase{notificationRepo: notificationRepo, clientRepo: clientRepo}
}

func (uc *PostNotificationUseCase) Execute(notification entities.Notification) error {
	client, err := uc.clientRepo.Search(notification.ClientID)
	if err != nil {
		return err
	}

	return uc.notificationRepo.Send(notification, client)
}
