package services

import "tienda/src/Notification/Application/repositories"

type RabbitMQService struct {
	rabbitMQUseCase *repositories.RabbitMQUseCase
}

func NewRabbitMQService(rabbitMQUseCase *repositories.RabbitMQUseCase) *RabbitMQService {
	return &RabbitMQService{rabbitMQUseCase: rabbitMQUseCase}
}

func (svc *RabbitMQService) SendMessage(message string, id int, name string) error {
	return svc.rabbitMQUseCase.Execute(message, id, name)
}

//socket.io
