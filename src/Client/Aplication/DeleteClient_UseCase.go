package application

import domain "tienda/src/Client/Domain"

type DeleteClientUseCase struct {
	repo domain.ClientRepository
}

func DeleteClient(repo domain.ClientRepository) *DeleteClientUseCase {
	return &DeleteClientUseCase{repo: repo}
}

func (uc *DeleteClientUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
