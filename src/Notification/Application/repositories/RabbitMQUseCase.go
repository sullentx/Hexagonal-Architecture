package repositories

import (
	domain "tienda/src/Notification/Domain"
)

type RabbitMQUseCase struct {
	repo domain.IMessagePublisher
}

func NewRabbitMQUseCase(repo domain.IMessagePublisher) *RabbitMQUseCase {
	return &RabbitMQUseCase{repo: repo}
}

func (uc *RabbitMQUseCase) Execute(message string, id int, name string) error {
	return uc.repo.PublishMessage(message, id, name)
}
