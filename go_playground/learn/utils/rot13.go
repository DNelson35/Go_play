package utils

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

// Read applies ROT13 encoding to the data read from the underlying reader.
//
// This method overrides the default Read method of the io.Reader interface.
// It reads data from the wrapped reader into the provided byte slice.
// For each byte read, it applies ROT13 encoding by shifting letters 13 places in the alphabet.
// Specifically, it shifts lowercase letters from 'a' to 'm' forward by 13 places,
// and from 'n' to 'z' backward by 13 places. Similarly, it shifts uppercase letters
// from 'A' to 'M' forward by 13 places, and from 'N' to 'Z' backward by 13 places.
// Non-alphabetic characters remain unchanged.
//
// Returns the number of bytes read and an error, if any occurred.
func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err == nil {
		for i := range b {
			switch true {
			case b[i] >= 97 && b[i] <= 109:
				b[i] += 13
			case b[i] > 109 && b[i] <= 122:
				b[i] -= 13
			case b[i] >= 65 && b[i] <= 77:
				b[i] += 13
			case b[i] > 77 && b[i] <= 90:
				b[i] -= 13
			}
		}
	}
	return n, err
}


// NewRot13Reader creates and returns a new rot13Reader instance.
//
// This wrapper type applies ROT13 encoding to the data read from the provided io.Reader.
// It is designed to be used with any object that satisfies the io.Reader interface,
// allowing for easy application of ROT13 encoding to various types of input streams,
// such as file content or network data.
//
// Parameters:
//   reader io.Reader: The underlying reader to apply ROT13 encoding to.
//
// Returns:
//   rot13Reader: A new rot13Reader instance wrapping the provided reader.
func NewRot13Reader(reader io.Reader) rot13Reader {
  return rot13Reader{r: reader}
}
