package keys_test

import (
	"slices"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/keys"
)

func TestKeys_Wrap_Unwrap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		passphrase string
		kGrain     []byte
		ad         []byte
		wantErr    bool
	}{
		{
			name:       "Test Case 1",
			passphrase: "password",
			kGrain:     hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ad:         []byte("test ad information"),
			wantErr:    false,
		},
		{
			name:       "Empty Passphrase",
			passphrase: "",
			kGrain:     hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ad:         []byte("test ad information"),
			wantErr:    false,
		},
		{
			name:       "Empty AD",
			passphrase: "password",
			kGrain:     hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ad:         []byte(""),
			wantErr:    false,
		},
		{
			name:       "Empty Passphrase and AD",
			passphrase: "",
			kGrain:     hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ad:         []byte(""),
			wantErr:    false,
		},
		{
			name:       "Long Passphrase",
			passphrase: `bcT5,enVo5-w4iotkrMY!Y,>+/AZH>w*y9Z,EL,2Vc5x.rn5*/EHdS4ivsEkYL_2-5:d/,kH33?vnmy3Ns!Hfud:\7;s?Fzk_pX9?/7E.Y2W;?uvDqRRUXDH\b!_/aiW`,
			kGrain:     hexDecode(t, "000102030405060708090a0b0c0d0e0f"),
			ad:         []byte("long passphrase test case"),
			wantErr:    false,
		},
		{
			name:       "Long kGrain",
			passphrase: `password`,
			kGrain:     hexDecode(t, "0000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f00102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f"),
			ad:         []byte("long kGrain test case"),
			wantErr:    true,
		},
		{
			name:       "Short kGrain",
			passphrase: `password`,
			kGrain:     hexDecode(t, "00010203040506070809"),
			ad:         []byte("long kGrain test case"),
			wantErr:    true,
		},
		{
			name:       "Long Passphrase and kGrain",
			passphrase: `iRtP_ucM<tiYzK-rMk,Fnjfdtb.V55k374JNy5\\y5:si5?Wi7*v7>3sD,h:RciMZvDqjPF<LyPd:Xf;zAorJHFvJCziUw/NX<;H5;Da\o5=TPT9R7ANhphcNhqF*-r/`,
			kGrain:     hexDecode(t, "0000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f00102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f"),
			ad:         []byte("long kGrain test case"),
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := keys.New()
			gotErr := k.Wrap(tt.passphrase, tt.kGrain, tt.ad)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Wrap() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Wrap() succeeded unexpectedly")
			}

			unwrappedKey, err := k.Unwrap(tt.passphrase, tt.ad)
			if err != nil {
				t.Errorf("Unwrap() failed: %v", err)
			}

			if !slices.Equal(unwrappedKey, tt.kGrain) {
				t.Errorf("Unwrap() = %v, want %v", unwrappedKey, tt.kGrain)
			}
		})
	}
}
