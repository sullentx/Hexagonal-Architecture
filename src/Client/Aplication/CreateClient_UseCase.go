package application

import (
	domain "tienda/src/client/domain"
	"tienda/src/client/domain/entities"
)

type CreateClientUseCase struct {
	clientRepo domain.ClientRepository
}

func NewCreateClientUseCase(clientRepo domain.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{clientRepo: clientRepo}
}

func (uc *CreateClientUseCase) Execute(client entities.Client) error {
	return uc.clientRepo.Create(client)
}
