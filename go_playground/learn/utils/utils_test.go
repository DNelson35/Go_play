package utils

import (
    "testing"
		"strings"
		"io"
)


// learn/utils/euler.go
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



// learn/utils/rot13.go
func TestRot13Reader(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Lowercase",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: "nopqrstuvwxyzabcdefghijklm",
		},
		{
			name:     "Uppercase",
			input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: "NOPQRSTUVWXYZABCDEFGHIJKLM",
		},
		{
			name:     "Mixed Case",
			input:    "Hello World!",
			expected: "Uryyb Jbeyq!",
		},
		{
			name:     "Non-Alphabetic",
			input:    "1234567890 !@#$%^&*()",
			expected: "1234567890 !@#$%^&*()", // No change expected
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// Convert the input string to an io.Reader
			reader := strings.NewReader(tt.input)

			// Create a new buffer to accumulate the decoded output
			var outBuf strings.Builder
			outWriter := &outBuf

			// Manually read from the rot13Reader and write to the output buffer
			_, err := io.Copy(outWriter, NewRot13Reader(reader))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Convert the accumulated output to a string for comparison
			actualOutput := outBuf.String()

			if actualOutput != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, actualOutput)
			}
		})
	}
}