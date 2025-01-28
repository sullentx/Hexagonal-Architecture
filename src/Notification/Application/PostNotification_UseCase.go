package application

import (
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type PostNotificationUseCase struct {
	repo domain.NotificationRepository
}

func NewPostNotification(repo domain.NotificationRepository) *PostNotificationUseCase {
	return &PostNotificationUseCase{repo: repo}
}

func (uc *PostNotificationUseCase) Execute(notification entities.Notification) error {
	return uc.repo.Send(notification)
}
