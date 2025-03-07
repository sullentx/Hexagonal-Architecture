package application

import (
	domain "tienda/src/client/domain"
	"tienda/src/client/domain/entities"
)

type ModifyClientUseCase struct {
	repo domain.ClientRepository
}

func NewModifyClient(repo domain.ClientRepository) *ModifyClientUseCase {
	return &ModifyClientUseCase{repo: repo}
}

func (uc *ModifyClientUseCase) Execute(id int, client entities.Client) error {
	return uc.repo.Modify(id, client)
}
