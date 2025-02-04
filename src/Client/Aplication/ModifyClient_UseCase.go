package application

import (
	domain "tienda/src/Client/Domain"
	"tienda/src/Client/Domain/entities"
)

type ModifyClientUseCase struct {
	repo domain.ClientRepository
}

func ModifyClient(repo domain.ClientRepository) *ModifyClientUseCase {
	return &ModifyClientUseCase{repo: repo}
}

func (uc *ModifyClientUseCase) Execute(id int, client entities.Client) error {
	return uc.repo.Modify(id, client)
}
