package domain

import (
	Client "tienda/src/client/domain/entities"
	entities "tienda/src/notification/domain/entities"
)

type NotificationRepository interface {
	Send(nofication entities.Notification, client Client.Client) error
	GetMessages() ([]entities.Notification, error)
	Search(id int) (entities.Notification, error)
	Delete(id int) error
	ModifyMessage(id int, content string) error
	GetNotifications(clientID int) ([]entities.Notification, error)
}
