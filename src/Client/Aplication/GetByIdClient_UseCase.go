package application

import (
	domain "tienda/src/client/domain"
	"tienda/src/client/domain/entities"
)

type GetByIdClientUseCase struct {
	repo domain.ClientRepository
}

func NewGetByIdClientUseCase(repo domain.ClientRepository) *GetByIdClientUseCase {
	return &GetByIdClientUseCase{repo: repo}
}

func (uc *GetByIdClientUseCase) Execute(id int) (entities.Client, error) {
	return uc.repo.Search(id)
}
