package applicationnegocio

import (
	domainnegocio "tienda/src/products/domain-negocio"
	"tienda/src/products/domain-negocio/entities"
)

type GetOneProductUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

// constructor
func GetOneProduct(repo domainnegocio.IproductrRepositoy) *GetOneProductUseCase {
	return &GetOneProductUseCase{repo: repo}
}

func (uc *GetOneProductUseCase) Execute(id int) (entities.Product, error) {
	return uc.repo.GetOne(id)
}
