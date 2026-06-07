package keys_test

import (
	"encoding/hex"
	"slices"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/keys"
)

func TestNewAES128CCM(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nonce   []byte
		ad      []byte
		want    *keys.AES128CCM
		wantErr bool
	}{
		{
			name:  "No Error",
			nonce: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			ad:    []byte("a test associate data"),
			want: &keys.AES128CCM{
				Nonce: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				AD:    []byte("a test associate data"),
			},
			wantErr: false,
		},
		{
			name:    "Empty Data",
			nonce:   []byte{},
			ad:      []byte{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Error Length",
			nonce:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			ad:      []byte{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := keys.NewAES128CCM(tt.nonce, tt.ad)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewAES128CCM() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewAES128CCM() succeeded unexpectedly")
			}

			if !slices.Equal(got.Nonce, tt.want.Nonce) ||
				!slices.Equal(got.AD, tt.want.AD) {
				t.Errorf("NewAES128CCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAES128CCM_Encrypt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		nonce []byte
		ad    []byte
		// Named input parameters for target function.
		key        []byte
		plaintext  []byte
		ciphertext []byte
		tag        []byte
		wantErr    bool
	}{
		{
			name:       "Test Case 1",
			nonce:      hexDecode(t, "000102030405060708090a0b"),
			ad:         []byte("test ad"),
			key:        hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			plaintext:  hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ciphertext: hexDecode(t, "3314f164d885c2b6791ac3eb0ee78b8f"),
			tag:        hexDecode(t, "55a8b5867228c73208fac44937cdd680"),
			wantErr:    false,
		},
		{
			name:       "Test Case 2",
			nonce:      hexDecode(t, "0c0b0a090807060504030201"),
			ad:         []byte{},
			key:        hexDecode(t, "0f0e0d0c0b0a09080706050403020100"),
			plaintext:  []byte("hello world!!!!!"),
			ciphertext: hexDecode(t, "1a2def10bf9bcbad7e9f4861ab1c1cf7"),
			tag:        hexDecode(t, "d09ca098733261dd0fa42ccdce8e99f1"),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := keys.NewAES128CCM(tt.nonce, tt.ad)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			ciphertext, tag, gotErr := a.Encrypt(tt.key, tt.plaintext)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Encrypt() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Encrypt() succeeded unexpectedly")
			}

			if !slices.Equal(ciphertext, tt.ciphertext) {
				t.Errorf("Encrypt() = %v, want %v", ciphertext, tt.ciphertext)
			}
			if !slices.Equal(tag, tt.tag) {
				t.Errorf("Encrypt() = %v, want %v", tag, tt.tag)
			}
		})
	}
}

func hexDecode(t *testing.T, s string) []byte {
	t.Helper()

	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("invalid hex %q: %v", s, err)
	}

	return b
}
