package utils

import "crypto/rand"

// Create a random byte slice with a given length.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
