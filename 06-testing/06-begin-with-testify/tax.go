package tax

import (
	"errors"
	"time"
)

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

func CalculateTaxSleep(amount float64) float64 {
	time.Sleep(1 * time.Millisecond)
	if amount == 0 {
		return 0.0
	}
	if amount <= 1000 {
		return 10.0
	}
	return 5.0
}
