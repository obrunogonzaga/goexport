package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmptyf(t, p.ID.String(), "ID should not be empty")
	assert.Equal(t, p.Name, "Product 1")
	assert.Equal(t, p.Price, 10.0)
}

func TestProductWhenNameIsRequered(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrorNameIsRequered)
}

func TestProductWhenPriceIsRequered(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrorPriceRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrInvalidPrice)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
