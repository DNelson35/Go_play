package leet

import (
	"testing"
	"strings"
	"io"
)

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