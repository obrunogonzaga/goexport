package entity

import (
	"errors"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/pkg/entity"
	"time"
)

var (
	ErrorIDIsRequered   = errors.New("ID is required")
	ErrorInvalidID      = errors.New("Invalid ID")
	ErrorNameIsRequered = errors.New("Name is required")
	ErrorPriceRequired  = errors.New("Price is required")
	ErrInvalidPrice     = errors.New("Invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequered
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameIsRequered
	}
	if p.Price == 0 {
		return ErrorPriceRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
