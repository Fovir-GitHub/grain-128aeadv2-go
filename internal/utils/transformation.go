package utils

import "fmt"

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
