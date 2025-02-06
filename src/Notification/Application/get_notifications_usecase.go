package application

import (
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type GetNotificationsUseCase struct {
	repo domain.NotificationRepository
}

func NewGetNotificationsUseCase(repo domain.NotificationRepository) *GetNotificationsUseCase {
	return &GetNotificationsUseCase{repo: repo}
}

func (uc *GetNotificationsUseCase) Execute(clientID int) ([]entities.Notification, error) {
	return uc.repo.GetNotifications(clientID)
}
