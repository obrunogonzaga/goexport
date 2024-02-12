package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 10.0, tax, "Tax for 1000 should be 10")

	tax, err = CalculateTax(0)
	assert.Error(t, err, "Amount cannot be zero")
	assert.Equal(t, 0.0, tax, "Tax for 0 should be 0")
	assert.Contains(t, err.Error(), "cannot be zero", "Error message should be 'Amount cannot be zero'")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("Save", 10.0).Return(nil).Once()
	repository.On("Save", 0.0).Return(errors.New("Error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err, "Error should be nil")

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "Error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 2)
}
