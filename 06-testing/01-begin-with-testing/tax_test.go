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
