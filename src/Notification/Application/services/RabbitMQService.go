package services

import "tienda/src/Notification/Application/repositories"

type RabbitMQService struct {
	rabbitMQUseCase *repositories.RabbitMQUseCase
}

func NewRabbitMQService(rabbitMQUseCase *repositories.RabbitMQUseCase) *RabbitMQService {
	return &RabbitMQService{rabbitMQUseCase: rabbitMQUseCase}
}

func (svc *RabbitMQService) SendMessage(message string, id int) error {
	return svc.rabbitMQUseCase.Execute(message, id)
}

//socket.io
