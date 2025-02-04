package application

import (
	domain "tienda/src/Client/Domain"
	"tienda/src/Client/Domain/entities"
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
