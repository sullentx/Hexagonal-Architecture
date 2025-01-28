package domain

import entities "tienda/src/Notification/Domain/Entities"

type NotificationRepository interface {
	Send(nofication entities.Notification) error
	GetMessages() ([]entities.Notification, error)
	Search(id int) (entities.Notification, error)
	Delete(id int) error
	ModifyMessage(id int, content string) error
}
