package application

import (
	domain "tienda/src/Notification/Domain"
	entities "tienda/src/Notification/Domain/Entities"
)

type SearchNotificationUseCase struct {
	repo domain.NotificationRepository
}

func SearchNotification(repo domain.NotificationRepository) *SearchNotificationUseCase {
	return &SearchNotificationUseCase{repo: repo}
}

func (uc *SearchNotificationUseCase) Execute(id int) (entities.Notification, error) {
	return uc.repo.Search(id)
}
