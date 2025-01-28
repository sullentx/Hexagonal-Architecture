package application

import domain "tienda/src/Notification/Domain"

type DeleteNotificationUseCase struct {
	repo domain.NotificationRepository
}

func DeleteNotification(repo domain.NotificationRepository) *DeleteNotificationUseCase {
	return &DeleteNotificationUseCase{repo: repo}
}

func (uc *DeleteNotificationUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
