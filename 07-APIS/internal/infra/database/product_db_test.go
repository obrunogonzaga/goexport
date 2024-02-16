package database

import (
	"fmt"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	db.AutoMigrate(entity.Product{})
	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
	assert.NotEmpty(t, products)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
	assert.NotEmpty(t, products)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
	assert.NotEmpty(t, products)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	db.AutoMigrate(entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	db.Create(product)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	db.AutoMigrate(entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	db.Create(product)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.Nil(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	db.AutoMigrate(entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	db.Create(product)
	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
}
