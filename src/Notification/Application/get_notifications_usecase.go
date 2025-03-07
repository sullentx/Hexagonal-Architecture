package application

import (
	domain "tienda/src/notification/domain"
	entities "tienda/src/notification/domain/entities"
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
