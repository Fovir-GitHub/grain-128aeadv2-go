package keys

import (
	"encoding/hex"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

func Test_kdf(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		passphrase string
		salt       string
		iter       int
		keysize    int
		want       string
		wantErr    bool
	}{
		{
			name:       "Test Case 1",
			passphrase: "password",
			salt:       "89eba7fd0141851a70919e4976e6cda8",
			iter:       600_000,
			keysize:    16,
			want:       "8329518cc01e8c822a13c35841d53227",
			wantErr:    false,
		},
		{
			name:       "Test Case 2",
			passphrase: "",
			salt:       "3f2dc3a3c0c964aeb3cb45a8fa2c9106",
			iter:       10,
			keysize:    16,
			want:       "7db0bb34ef975b339674bf1a40c4e5be",
			wantErr:    false,
		},
		{
			name:       "Long Passphrase",
			passphrase: "ni>TM?!=F9>4sT*=*2q3AcqHHDAjjDgKDgDvv?NLtApoSV9k!F5XTLm7pk<inaQDoej4uaAK*AK5UmtJQy!H93E4Ps?vy4vrdtvyLoFN9TH<n!NjRwjdVzyeEwhTEKmC",
			salt:       "d1efa2cda9f94895f683a8f3a2dbc798",
			iter:       3200,
			keysize:    16,
			want:       "e174241b63caf378fd18b94e4762fc2d",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salt, _ := utils.Hex2Byte(tt.salt)
			got, gotErr := kdf(tt.passphrase, salt, tt.iter, tt.keysize)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("KDF() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("KDF() succeeded unexpectedly")
			}

			gotStr := hex.EncodeToString(got)
			if tt.want != gotStr {
				t.Errorf("KDF() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}
