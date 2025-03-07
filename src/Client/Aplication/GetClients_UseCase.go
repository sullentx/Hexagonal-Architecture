package application

import (
	domain "tienda/src/client/domain"
	"tienda/src/client/domain/entities"
)

type GetClientsUseCase struct {
	repo domain.ClientRepository
}

func NewGetClients(repo domain.ClientRepository) *GetClientsUseCase {
	return &GetClientsUseCase{repo: repo}
}

func (uc *GetClientsUseCase) Execute() ([]entities.Client, error) {
	return uc.repo.GetClients()
}
