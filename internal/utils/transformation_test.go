package utils_test

import (
	"slices"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

func TestByte2Bits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		b    []byte
		want []int
	}{
		{
			name: "Single Byte Test",
			b:    []byte("a"),
			want: []int{1, 0, 0, 0, 0, 1, 1, 0},
		},
		{
			name: "Multiple Byte Test",
			b:    []byte("abcdefg"),
			want: []int{1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.Byte2Bits(tt.b)
			if !slices.Equal(tt.want, got) {
				t.Errorf("Byte2Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits2Byte(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bits    []int
		want    []byte
		wantErr bool
	}{
		{
			name:    "Error Test",
			bits:    []int{1, 2, 3, 4, 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Single Byte Test",
			want:    []byte("a"),
			bits:    []int{1, 0, 0, 0, 0, 1, 1, 0},
			wantErr: false,
		},
		{
			name:    "Multiple Byte Test",
			want:    []byte("abcdefg"),
			bits:    []int{1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.Bits2Byte(tt.bits)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Bits2Byte() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Bits2Byte() succeeded unexpectedly")
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Bits2Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHex2Byte(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Empty String",
			s:       "",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "Only Prefix",
			s:       "0x",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "Single Byte",
			s:       "0x61",
			want:    []byte{0x61},
			wantErr: false,
		},
		{
			name:    "Uppercase Prefix",
			s:       "0X61",
			want:    []byte{0x61},
			wantErr: false,
		},
		{
			name:    "Multiple Bytes",
			s:       "0x123456",
			want:    []byte{0x12, 0x34, 0x56},
			wantErr: false,
		},
		{
			name:    "Mixed Case",
			s:       "0xDeAdBeEf",
			want:    []byte{0xde, 0xad, 0xbe, 0xef},
			wantErr: false,
		},
		{
			name:    "Grain Key Example",
			s:       "000102030405060708090a0b0c0d0e0f",
			want:    []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			wantErr: false,
		},
		{
			name:    "Odd Length",
			s:       "0x123",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid Character",
			s:       "0x12gg",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid Prefix Content",
			s:       "0xz123",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.Hex2Byte(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Hex2Byte() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Hex2Byte() succeeded unexpectedly")
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Hex2Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}
