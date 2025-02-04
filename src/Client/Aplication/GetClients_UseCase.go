package application

import (
	domain "tienda/src/Client/Domain"
	"tienda/src/Client/Domain/entities"
)

type GetClientsUseCase struct {
	repo domain.ClientRepository
}

func GetClients(repo domain.ClientRepository) *GetClientsUseCase {
	return &GetClientsUseCase{repo: repo}
}

func (uc *GetClientsUseCase) Execute() ([]entities.Client, error) {
	return uc.repo.GetClients()
}
