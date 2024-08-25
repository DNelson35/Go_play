package euler

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

func TestPrimeFactor(t *testing.T){
testData := []struct {
	name string
	input int
	expected int
}{
	{
		name: "test 1",
		input: 600851475143,
		expected: 6857,
	},
	{
		name: "test 2",
		input: 0,
		expected: 0,
	},
	{
		name: "test 3",
		input: 14,
		expected: 7,
	},
	{
		name: "test 4",
		input: 25,
		expected: 5,
	},
}

for _, test := range testData {
	t.Run(test.name, func(t *testing.T) {
		r := Primefactor(test.input)

		if r != test.expected {
			t.Errorf("Expected: %d got %d", test.expected, r )
		}
	})
}
}

func TestLgPalindrom(t *testing.T) {
expected := 906609 
result := LgPalindrom()

if result != expected {
	t.Errorf("LgPalindrom() = %d; want %d", result, expected)
}
}

func TestFindDiffSquares(t *testing.T) {
tests := []struct {
	x        int
	expected int
}{
	{1, 0},     
	{2, 4},     
	{3, 22},    
	{10, 2640},      
	{0, 0},          
}

for _, tt := range tests {
	result := FindDiffSquares(tt.x)
	if result != tt.expected {
		t.Errorf("FindDiffSquares(%d) = %d; want %d", tt.x, result, tt.expected)
	}
}
}

func TestGetPrime(t *testing.T) {
tests := []struct {
		n        int
		expected int
}{
		{1, 2},   
		{2, 3},   
		{3, 5},   
		{4, 7},   
		{5, 11},  
		{10, 29},  
		{20, 71},  
		{0, 0},     
		{-5, 0},     
}

for _, tt := range tests {
	result := GetPrime(tt.n)
	if result != tt.expected {
		t.Errorf("GetPrime(%d) = %d; want %d", tt.n, result, tt.expected)
	}
}
}