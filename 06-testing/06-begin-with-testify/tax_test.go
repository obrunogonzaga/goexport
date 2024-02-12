package tax

import (
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
