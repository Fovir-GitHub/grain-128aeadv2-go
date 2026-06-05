package grain

import (
	"slices"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

func TestGrain128AEADV2_Encrypt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		key   string
		nonce string
		// Named input parameters for target function.
		plaintext [][]byte
	}{
		{
			name:  "Test Vector 1",
			key:   "0x00000000000000000000000000000000",
			nonce: "0x000000000000000000000000",
			plaintext: [][]byte{
				[]byte("This is a top-secret message. Your implementation should be able to handle variable length input messages, and the inputs should not be hardcoded."),
				[]byte("Test plaintext for vector 1"),
			},
		},
		{
			name:  "Test Vector 2",
			key:   "0x000102030405060708090a0b0c0d0e0f",
			nonce: "0x000102030405060708090a0b",
			plaintext: [][]byte{
				[]byte("This is not a top-secret message."),
				[]byte("Test plaintext for vector 2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, _ := utils.Hex2Byte(tt.key)
			n, _ := utils.Hex2Byte(tt.nonce)
			g, err := New(k, n)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}

			for _, p1 := range tt.plaintext {
				c := g.Encrypt(p1)
				p2 := g.Decrypt(c)
				if !slices.Equal(p1, p2) {
					t.Errorf("p1 = %v, p2 = %v", p1, p2)
				}
			}
		})
	}
}
