package keys_test

import (
	"bytes"
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

func TestAES128CCM_Auth(t *testing.T) {
	mkEncryption := func(key, nonce, plaintext, ad []byte) (ct, tag []byte) {
		a, err := keys.NewAES128CCM(nonce, ad)
		if err != nil {
			t.Fatalf("mkEncryption: NewAES128CCM failed (nonce=%v, ad=%v): %v", nonce, ad, err)
		}
		ct, tag, err = a.Encrypt(key, plaintext)
		if err != nil {
			t.Fatalf("mkEncryption: Encrypt failed: %v", err)
		}
		return ct, tag
	}

	validKey := []byte("0123456789abcdef")
	validNonce := []byte("123456789012")
	validAD := []byte("additional data")
	validPlain := []byte("hello world")

	validCT, validTag := mkEncryption(validKey, validNonce, validPlain, validAD)

	tamperedTag := func() []byte {
		b := make([]byte, len(validTag))
		copy(b, validTag)
		b[0] ^= 0xFF
		return b
	}()
	tamperedCT := func() []byte {
		b := make([]byte, len(validCT))
		copy(b, validCT)
		b[0] ^= 0xFF
		return b
	}()

	tests := []struct {
		name       string
		nonce      []byte
		ad         []byte
		key        []byte
		ciphertext []byte
		tag        []byte
		want       []byte
		wantErr    bool
	}{
		{
			name:       "Valid: Correct key/nonce/ad/tag",
			nonce:      validNonce,
			ad:         validAD,
			key:        validKey,
			ciphertext: validCT,
			tag:        validTag,
			want:       validPlain,
			wantErr:    false,
		},
		{
			name:       "Error: Tampered Tag",
			nonce:      validNonce,
			ad:         validAD,
			key:        validKey,
			ciphertext: validCT,
			tag:        tamperedTag,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Error: Tampered Ciphertext",
			nonce:      validNonce,
			ad:         validAD,
			key:        validKey,
			ciphertext: tamperedCT,
			tag:        validTag,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Error: Wrong AD",
			nonce:      validNonce,
			ad:         []byte("wrong ad"),
			key:        validKey,
			ciphertext: validCT,
			tag:        validTag,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Error: Invalid Key Length",
			nonce:      validNonce,
			ad:         validAD,
			key:        []byte("shortkey"),
			ciphertext: validCT,
			tag:        validTag,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Valid: Empty Plaintext",
			nonce:      validNonce,
			ad:         validAD,
			key:        validKey,
			ciphertext: func() []byte { ct, _ := mkEncryption(validKey, validNonce, []byte{}, validAD); return ct }(),
			tag:        func() []byte { _, tag := mkEncryption(validKey, validNonce, []byte{}, validAD); return tag }(),
			want:       []byte{},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := keys.NewAES128CCM(tt.nonce, tt.ad)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got, gotErr := a.Auth(tt.key, tt.ciphertext, tt.tag)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Auth() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Auth() succeeded unexpectedly")
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
