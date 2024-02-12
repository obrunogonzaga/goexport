package tax

import (
	"errors"
)

type Repository interface {
	Save(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTaxMock(amount)
	return repository.Save(tax)
}

func CalculateTax(amount float64) (float64, error) {
	if amount < 0 {
		return 0.0, errors.New("Amount cannot be negative")
	}
	if amount == 0 {
		return 0.0, errors.New("Amount cannot be zero")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

func CalculateTaxMock(amount float64) float64 {
	if amount < 0 {
		return 0.0
	}
	if amount == 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}
