package applicationnegocio

import (
	domainnegocio "tienda/src/products/domain-negocio"
	"tienda/src/products/domain-negocio/entities"
)

type GetNewProductsUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

func NewGetNewProductsUseCase(repo domainnegocio.IproductrRepositoy) *GetNewProductsUseCase {
	return &GetNewProductsUseCase{repo: repo}
}

func (uc *GetNewProductsUseCase) Execute(lastProductID int) ([]entities.Product, int, error) {
	products, latestID, err := uc.repo.GetNewProducts(lastProductID)
	if err != nil {
		return nil, lastProductID, err
	}
	return products, latestID, nil
}
