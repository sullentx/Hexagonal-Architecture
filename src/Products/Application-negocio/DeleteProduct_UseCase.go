package applicationnegocio

import domainnegocio "tienda/src/Products/Domain-negocio"

type DeleteProductUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

func DeleteProduct(repo domainnegocio.IproductrRepositoy) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
