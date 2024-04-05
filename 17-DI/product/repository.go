package product

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) GetProductByID(id int) (Product, error) {
	return Product{
		ID:   1,
		Name: "Product 1",
	}, nil
}
