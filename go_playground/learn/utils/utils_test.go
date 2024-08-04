package utils

import (
    "testing"
)

func TestFibonacci(t *testing.T) {
    f := Fibonacci()
	for i := 0; i < 10; i++ {
		num := f(i)
		if i == 1 && num != 1 {
			t.Errorf("Expected 1, got %d", num)
		}

		if i == 9 && num != 34 {
			t.Errorf("Expected 34, got %d", num)
		}
	}
  
}

func TestSumMul(t *testing.T) {
	sum := SumMul(3, 5)
	if sum != 233168{
		t.Errorf("Expected 233168, got %d", sum)
	}
}