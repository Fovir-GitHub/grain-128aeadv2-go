package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Byte2Bits transforms byte to bits format (LSB-First).
func Byte2Bits(b []byte) []int {
	bits := make([]int, len(b)*8)
	for i, v := range b {
		for j := range 8 {
			bits[i*8+j] = ((int(v) >> j) & 1)
		}
	}

	return bits
}

// Bits2Byte transforms bits (LSB-First) to byte format.
func Bits2Byte(bits []int) ([]byte, error) {
	if len(bits)%8 != 0 {
		return nil, fmt.Errorf("the length of bits is not divided by 8")
	}

	if err := validateBits(bits); err != nil {
		return nil, err
	}

	bytes := make([]byte, len(bits)/8)
	for i := range bytes {
		b := 0
		for j := range 8 {
			b |= bits[i*8+j] << j
		}
		bytes[i] = byte(b)
	}
	return bytes, nil
}

// Hex2Byte trims the prefix `0x` in a hex string, and convert it to bytes.
func Hex2Byte(s string) ([]byte, error) {
	s = strings.TrimPrefix(strings.ToLower(s), "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("decode hex string %s failed: %w", s, err)
	}
	return b, nil
}

// Bits2Hex transforms bits to hex string.
func Bits2Hex(bits []int) (string, error) {
	b, err := Bits2Byte(bits)
	if err != nil {
		return "", fmt.Errorf("bits to hex failed: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// validateBits checks whether a bit is valid (0 or 1).
func validateBits(bits []int) error {
	for _, b := range bits {
		if b != 0 && b != 1 {
			return fmt.Errorf("invalid bit: %v", b)
		}
	}
	return nil
}
