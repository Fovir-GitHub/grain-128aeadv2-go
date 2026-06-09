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

			gotLFSR, gotNFSR := g.loadLFSRNFSR()

			if gotLFSR != tt.wantLFSR {
				t.Errorf("gotLFSR = %v, want %v", gotLFSR, tt.wantLFSR)
			}

			if gotNFSR != tt.wantNFSR {
				t.Errorf("gotNFSR = %v, want %v", gotNFSR, tt.wantNFSR)
			}
		})
	}
}

func TestGrain128AEADV2_Init(t *testing.T) {
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
			wantLFSR: "8f395a9421b0963364e2ed30679c8ee1",
			wantNFSR: "81f7e0c655d035823310c278438dbc20",
		},
		{
			name:     "Test Vector 2",
			key:      "0x000102030405060708090a0b0c0d0e0f",
			nonce:    "0x000102030405060708090a0b",
			wantLFSR: "0e1f950d45e05087c4cd63fd00eab310",
			wantNFSR: "b3c2e1b1eec1f08c2d6eae957f6af9d0",
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
			_, _, gotLFSR, gotNFSR := g.Init()

			if gotLFSR != tt.wantLFSR {
				t.Errorf("gotLFSR = %v, want %v", gotLFSR, tt.wantLFSR)
			}

			if gotNFSR != tt.wantNFSR {
				t.Errorf("gotNFSR = %v, want %v", gotNFSR, tt.wantNFSR)
			}
		})
	}
}
