package application

import domain "tienda/src/Notification/Domain"

type ModifyNotificationUseCase struct {
	repo domain.NotificationRepository
}

func ModifyNotification(repo domain.NotificationRepository) *ModifyNotificationUseCase {
	return &ModifyNotificationUseCase{repo: repo}
}

func (uc *ModifyNotificationUseCase) Execute(id int, message string) error {
	return uc.repo.ModifyMessage(id, message)
}
