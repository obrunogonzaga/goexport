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
		{1001, 10},
		{500, 5},
		{5000, 10},
		{0, 0},
	}
	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

func BenchmarkCalculateTaxSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxSleep(1000)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2, 5, 500.0, 1000.0, 1501}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0, got %f", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Expected 20, got %f", result)
		}
	})
}
