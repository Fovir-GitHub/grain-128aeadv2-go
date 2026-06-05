package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func Byte2Bits(b []byte) []int {
	bits := make([]int, len(b)*8)
	for i, v := range b {
		for j := range 8 {
			bits[i*8+j] = ((int(v) >> j) & 1)
		}
	}

	return bits
}

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

func Hex2Byte(s string) ([]byte, error) {
	s = strings.TrimPrefix(strings.ToLower(s), "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("decode hex string %s failed: %w", s, err)
	}
	return b, nil
}

func Bits2Hex(bits []int) (string, error) {
	b, err := Bits2Byte(bits)
	if err != nil {
		return "", fmt.Errorf("bits to hex failed: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func validateBits(bits []int) error {
	for _, b := range bits {
		if b != 0 && b != 1 {
			return fmt.Errorf("invalid bit: %v", b)
		}
	}
	return nil
}
