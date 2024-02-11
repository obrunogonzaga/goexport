package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 1000.0
	expected := 10.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{1000, 10},
		{1001, 5},
		{500, 10},
		{5000, 5},
	}
	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}
