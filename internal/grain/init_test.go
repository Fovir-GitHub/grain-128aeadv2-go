package grain

import (
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

func TestGrain128AEADV2_initLFSRNFSR(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		key      string
		nonce    string
		wantLFSR string
		wantNFSR string
	}{
		{
			name:     "Test Vector 1",
			key:      "0x00000000000000000000000000000000",
			nonce:    "0x000000000000000000000000",
			wantLFSR: "000000000000000000000000ffffff7f",
			wantNFSR: "00000000000000000000000000000000",
		},
		{
			name:     "Test Vector 2",
			key:      "0x000102030405060708090a0b0c0d0e0f",
			nonce:    "0x000102030405060708090a0b",
			wantLFSR: "000102030405060708090a0bffffff7f",
			wantNFSR: "000102030405060708090a0b0c0d0e0f",
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
			g.initLFSRNFSR()

			gotLFSR, err := utils.Bits2Hex(g.LFSR)
			if err != nil {
				t.Fatalf("could not convert LFSR to hex: %v", err)
			}
			if gotLFSR != tt.wantLFSR {
				t.Errorf("gotLFSR = %v, want %v", gotLFSR, tt.wantLFSR)
			}

			gotNFSR, err := utils.Bits2Hex(g.NFSR)
			if err != nil {
				t.Fatalf("could not convert NFSR to hex: %v", err)
			}
			if gotNFSR != tt.wantNFSR {
				t.Errorf("gotNFSR = %v, want %v", gotNFSR, tt.wantNFSR)
			}
		})
	}
}
