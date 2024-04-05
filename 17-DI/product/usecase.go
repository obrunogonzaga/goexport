package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

func (pu *ProductUseCase) GetProductByID(id int) (Product, error) {
	return pu.repository.GetProductByID(id)
}
