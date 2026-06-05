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
