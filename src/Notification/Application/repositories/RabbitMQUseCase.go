package repositories

import "tienda/src/notification/domain"

type RabbitMQUseCase struct {
	repo domain.IMessagePublisher
}

func NewRabbitMQUseCase(repo domain.IMessagePublisher) *RabbitMQUseCase {
	return &RabbitMQUseCase{repo: repo}
}

func (uc *RabbitMQUseCase) Execute(message string, id int, name string) error {
	return uc.repo.PublishMessage(message, id, name)
}
