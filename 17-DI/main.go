package main

import (
	"database/sql"
	"fmt"
	"github.com/obrunogonzaga/pos-go-expert/17-DI/product"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	useCase := product.NewProductUseCase(repository)

	product, err := useCase.GetProductByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
