package utils

import (
    "testing"
		"strings"
		"io"
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