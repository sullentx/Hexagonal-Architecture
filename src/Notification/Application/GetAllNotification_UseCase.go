package application

import (
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type GetAllNotificationUseCase struct {
	repo domain.NotificationRepository
}

func GetAllNotification(repo domain.NotificationRepository) *GetAllNotificationUseCase {
	return &GetAllNotificationUseCase{repo: repo}
}

func (uc *GetAllNotificationUseCase) Execute() ([]entities.Notification, error) {
	return uc.repo.GetMessages()
}
