package service_test

import (
	"encoding/hex"
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/service"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

func TestCipherService_Encrypt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *model.EncryptionRequest
		want    *model.EncryptionResp
		wantErr bool
	}{
		{
			name: "ASCII Plaintext",
			req: &model.EncryptionRequest{
				Plaintext:      "This is a top-secret message. Your implementation should be able to handle variable length input messages, and the inputs should not be hardcoded.",
				IsPlaintextHex: false,
				Nonce:          "0x000000000000000000000000",
				Key:            "0x00000000000000000000000000000000",
			},
			want: &model.EncryptionResp{
				Ciphertext: "00000000000000000000000063e8825e421ddf0c8b946503f5cd614e71a3ab1aa541f8602685e7e561549a3c243bdcaa41d89a493806e31e8b4a1aea28a7cd38be4d11e4ac9be82f7cfb9bb0509e3f0a5c8da188b957ca065a16ec30962ac25608f6b6c66fd359134a55aff5944daee2a53b3a6f45d296b64cd1112c1d84673abc252e819eda441c195427f6f1dff47e633b14cb41d961765cbc895c0274",
				LoadedLFSR: "000000000000000000000000ffffff7f",
				LoadedNFSR: "00000000000000000000000000000000",
				InitLFSR:   "8f395a9421b0963364e2ed30679c8ee1",
				InitNFSR:   "81f7e0c655d035823310c278438dbc20",
			},
			wantErr: false,
		},
		{
			name: "Hex Plaintext",
			req: &model.EncryptionRequest{
				Plaintext:      "54686973206973206120746f702d736563726574206d6573736167652e20596f757220696d706c656d656e746174696f6e2073686f756c642062652061626c6520746f2068616e646c65207661726961626c65206c656e67746820696e707574206d657373616765732c20616e642074686520696e707574732073686f756c64206e6f742062652068617264636f6465642e",
				IsPlaintextHex: true,
				Nonce:          "0x000000000000000000000000",
				Key:            "0x00000000000000000000000000000000",
			},
			want: &model.EncryptionResp{
				Ciphertext: "00000000000000000000000063e8825e421ddf0c8b946503f5cd614e71a3ab1aa541f8602685e7e561549a3c243bdcaa41d89a493806e31e8b4a1aea28a7cd38be4d11e4ac9be82f7cfb9bb0509e3f0a5c8da188b957ca065a16ec30962ac25608f6b6c66fd359134a55aff5944daee2a53b3a6f45d296b64cd1112c1d84673abc252e819eda441c195427f6f1dff47e633b14cb41d961765cbc895c0274",
				LoadedLFSR: "000000000000000000000000ffffff7f",
				LoadedNFSR: "00000000000000000000000000000000",
				InitLFSR:   "8f395a9421b0963364e2ed30679c8ee1",
				InitNFSR:   "81f7e0c655d035823310c278438dbc20",
			},
			wantErr: false,
		},
		{
			name: "Hex Plaintext with Prefix",
			req: &model.EncryptionRequest{
				Plaintext:      "0x54686973206973206120746f702d736563726574206d6573736167652e20596f757220696d706c656d656e746174696f6e2073686f756c642062652061626c6520746f2068616e646c65207661726961626c65206c656e67746820696e707574206d657373616765732c20616e642074686520696e707574732073686f756c64206e6f742062652068617264636f6465642e",
				IsPlaintextHex: true,
				Nonce:          "0x000000000000000000000000",
				Key:            "0x00000000000000000000000000000000",
			},
			want: &model.EncryptionResp{
				Ciphertext: "00000000000000000000000063e8825e421ddf0c8b946503f5cd614e71a3ab1aa541f8602685e7e561549a3c243bdcaa41d89a493806e31e8b4a1aea28a7cd38be4d11e4ac9be82f7cfb9bb0509e3f0a5c8da188b957ca065a16ec30962ac25608f6b6c66fd359134a55aff5944daee2a53b3a6f45d296b64cd1112c1d84673abc252e819eda441c195427f6f1dff47e633b14cb41d961765cbc895c0274",
				LoadedLFSR: "000000000000000000000000ffffff7f",
				LoadedNFSR: "00000000000000000000000000000000",
				InitLFSR:   "8f395a9421b0963364e2ed30679c8ee1",
				InitNFSR:   "81f7e0c655d035823310c278438dbc20",
			},
			wantErr: false,
		},
		{
			name: "Nonce and Key without Prefix",
			req: &model.EncryptionRequest{
				Plaintext:      "54686973206973206120746f702d736563726574206d6573736167652e20596f757220696d706c656d656e746174696f6e2073686f756c642062652061626c6520746f2068616e646c65207661726961626c65206c656e67746820696e707574206d657373616765732c20616e642074686520696e707574732073686f756c64206e6f742062652068617264636f6465642e",
				IsPlaintextHex: true,
				Nonce:          "000000000000000000000000",
				Key:            "00000000000000000000000000000000",
			},
			want: &model.EncryptionResp{
				Ciphertext: "00000000000000000000000063e8825e421ddf0c8b946503f5cd614e71a3ab1aa541f8602685e7e561549a3c243bdcaa41d89a493806e31e8b4a1aea28a7cd38be4d11e4ac9be82f7cfb9bb0509e3f0a5c8da188b957ca065a16ec30962ac25608f6b6c66fd359134a55aff5944daee2a53b3a6f45d296b64cd1112c1d84673abc252e819eda441c195427f6f1dff47e633b14cb41d961765cbc895c0274",
				LoadedLFSR: "000000000000000000000000ffffff7f",
				LoadedNFSR: "00000000000000000000000000000000",
				InitLFSR:   "8f395a9421b0963364e2ed30679c8ee1",
				InitNFSR:   "81f7e0c655d035823310c278438dbc20",
			},
			wantErr: false,
		},
		{
			name: "ASCII Plaintext 2",
			req: &model.EncryptionRequest{
				Plaintext:      "Test plaintext for vector 2",
				IsPlaintextHex: false,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Hex Plaintext 2",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Empty Key and Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "",
				Key:            "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Plaintext",
			req: &model.EncryptionRequest{
				Plaintext:      "",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0b",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Empty Key",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Hex",
			req: &model.EncryptionRequest{
				Plaintext:      "Test plaintext for vector 2",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Hex String As ASCII Plaintext",
			req: &model.EncryptionRequest{
				Plaintext:      "353436353733373432303730366336313639366537343635373837343230363636663732323037363635363337343666373232303332",
				IsPlaintextHex: false,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bd91455b136c9bc8502a1e38c979d27c28eb28709d2bda5bd9a488975f5698072d9e13fcf6e7a0afccd9dfb07ff23be67c08d97814538bf8575d1de92ff709345bb873c0e5f2ea90374ee6c4ff7cb193709fab6e2f671f3e257d1719527ea8be2874e1a21d938c837d07d15e3",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Invalid Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "Test plaintext for vector 2",
				IsPlaintextHex: false,
				Nonce:          "0x000102030405060708090a0babcd",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Key",
			req: &model.EncryptionRequest{
				Plaintext:      "Test plaintext for vector 2",
				IsPlaintextHex: false,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090a0b0c0d0e0fa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Key and Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "Test plaintext for vector 2",
				IsPlaintextHex: false,
				Nonce:          "0x000102030405060708090a0ba",
				Key:            "0x000102030405060708090a0b0c0d0e0fa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uppercase Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090A0B",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Mix-case Nonce",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090A0b",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Nonce with Surrounding Spaces",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "  0x000102030405060708090a0b  ",
				Key:            "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Uppercase Key",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090A0B0C0D0E0F",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Mix-case Key",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "0x000102030405060708090A0b0c0D0E0F",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Key with Surrounding Spaces",
			req: &model.EncryptionRequest{
				Plaintext:      "5465737420706c61696e7465787420666f7220766563746f722032",
				IsPlaintextHex: true,
				Nonce:          "0x000102030405060708090a0b",
				Key:            "  0x000102030405060708090a0b0c0d0e0f  ",
			},
			want: &model.EncryptionResp{
				Ciphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e service.CipherService
			got, gotErr := e.Encrypt(tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Encrypt() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Encrypt() succeeded unexpectedly")
			}

			checnField := func(name string, got, want any) {
				if got != want {
					t.Errorf("%s: Encrypt() = %v, want %v", name, got, want)
				}
			}

			checnField("Ciphertext", got.Ciphertext, tt.want.Ciphertext)
			checnField("LoadLFSR", got.LoadedLFSR, tt.want.LoadedLFSR)
			checnField("LoadNFSR", got.LoadedNFSR, tt.want.LoadedNFSR)
			checnField("InitLFSR", got.InitLFSR, tt.want.InitLFSR)
			checnField("InitNFSR", got.InitNFSR, tt.want.InitNFSR)
		})
	}
}

func TestCipherService_Decrypt(t *testing.T) {
	tests := []struct {
		name    string
		req     *model.DecryptionRequest
		want    *model.DecryptionResp
		wantErr bool
	}{
		{
			name: "Decrypt Vector 1",
			req: &model.DecryptionRequest{
				NonceCiphertext: "00000000000000000000000063e8825e421ddf0c8b946503f5cd614e71a3ab1aa541f8602685e7e561549a3c243bdcaa41d89a493806e31e8b4a1aea28a7cd38be4d11e4ac9be82f7cfb9bb0509e3f0a5c8da188b957ca065a16ec30962ac25608f6b6c66fd359134a55aff5944daee2a53b3a6f45d296b64cd1112c1d84673abc252e819eda441c195427f6f1dff47e633b14cb41d961765cbc895c0274",
				Key:             "0x00000000000000000000000000000000",
			},
			want: &model.DecryptionResp{
				Plaintext:  "54686973206973206120746f702d736563726574206d6573736167652e20596f757220696d706c656d656e746174696f6e2073686f756c642062652061626c6520746f2068616e646c65207661726961626c65206c656e67746820696e707574206d657373616765732c20616e642074686520696e707574732073686f756c64206e6f742062652068617264636f6465642e",
				LoadedLFSR: "000000000000000000000000ffffff7f",
				LoadedNFSR: "00000000000000000000000000000000",
				InitLFSR:   "8f395a9421b0963364e2ed30679c8ee1",
				InitNFSR:   "81f7e0c655d035823310c278438dbc20",
			},
			wantErr: false,
		},
		{
			name: "Decrypt Vector 2",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Decrypt Empty Plaintext",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0b",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Key without Prefix",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "000102030405060708090a0b0c0d0e0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Uppercase Key",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "0x000102030405060708090A0B0C0D0E0F",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Mix-case Key",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "0x000102030405060708090A0b0C0d0E0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Key with Surrounding Spaces",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "  0x000102030405060708090a0b0c0d0e0f  ",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Ciphertext with Surrounding Spaces",
			req: &model.DecryptionRequest{
				NonceCiphertext: "  000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d  ",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "5465737420706c61696e7465787420666f7220766563746f722032",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Empty Key",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Ciphertext",
			req: &model.DecryptionRequest{
				NonceCiphertext: "",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Key and Ciphertext",
			req: &model.DecryptionRequest{
				NonceCiphertext: "",
				Key:             "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Ciphertext Too Short (only nonce, no ciphertext)",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0b",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want: &model.DecryptionResp{
				Plaintext:  "",
				LoadedLFSR: "000102030405060708090a0bffffff7f",
				LoadedNFSR: "000102030405060708090a0b0c0d0e0f",
				InitLFSR:   "0e1f950d45e05087c4cd63fd00eab310",
				InitNFSR:   "b3c2e1b1eec1f08c2d6eae957f6af9d0",
			},
			wantErr: false,
		},
		{
			name: "Ciphertext Too Short (less than nonce length)",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Hex Ciphertext",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bgg4415f1",
				Key:             "0x000102030405060708090a0b0c0d0e0f",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Hex Key",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "0x000102030405060708090a0b0c0d0e0g",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Key Length",
			req: &model.DecryptionRequest{
				NonceCiphertext: "000102030405060708090a0bbe4415f1258fe3d158f8a4dadcde3490d2f2944f84e9e2e2db5e8d",
				Key:             "0x000102030405060708090a0b0c0d0e0fab",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c service.CipherService
			got, gotErr := c.Decrypt(tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Decrypt() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Decrypt() succeeded unexpectedly")
			}

			checkField := func(name string, got, want any) {
				if got != want {
					t.Errorf("%s: Decrypt() = %v, want %v", name, got, want)
				}
			}

			checkField("Plaintext", got.Plaintext, tt.want.Plaintext)
			checkField("LoadedLFSR", got.LoadedLFSR, tt.want.LoadedLFSR)
			checkField("LoadedNFSR", got.LoadedNFSR, tt.want.LoadedNFSR)
			checkField("InitLFSR", got.InitLFSR, tt.want.InitLFSR)
			checkField("InitNFSR", got.InitNFSR, tt.want.InitNFSR)
		})
	}
}

// TestService_EncryptionFunctionality implements Test Case 1:
// Encrypt P with (K,V) -> C, (K',V) -> C', (K,V') -> C”.
// All three ciphertexts (hex(nonce||cipher)) must differ.
func TestService_EncryptionFunctionality(t *testing.T) {
	var svc service.CipherService

	plaintext := "0x1234"
	k1 := "0x00000000000000000000000000000000"
	k2 := "0x000102030405060708090a0b0c0d0e0f"
	n1 := "0x000000000000000000000000"
	n2 := "0x000102030405060708090a0b"

	enc := func(key, nonce string) string {
		resp, err := svc.Encrypt(&model.EncryptionRequest{
			Plaintext: plaintext, IsPlaintextHex: true,
			Nonce: nonce, Key: key,
		})
		if err != nil {
			t.Fatalf("Encrypt() failed: %v", err)
		}
		return resp.Ciphertext
	}

	c1 := enc(k1, n1)
	t.Logf("1a) plaintext=%s, key=%s, nonce=%s, ciphertext=%s", plaintext, k1, n1, c1)

	c2 := enc(k2, n1)
	t.Logf("1b) plaintext=%s, key=%s, nonce=%s, ciphertext=%s", plaintext, k2, n1, c2)

	c3 := enc(k1, n2)
	t.Logf("1c) plaintext=%s, key=%s, nonce=%s, ciphertext=%s", plaintext, k1, n2, c3)

	if c1 == c2 || c1 == c3 || c2 == c3 {
		t.Error("ciphertexts should all differ when key or IV differs")
	}
}

// TestService_DecryptionFunctionality implements Test Case 2:
// Decrypt C with (K,V), C' with (K',V), C” with (K,V') to recover P.
func TestService_DecryptionFunctionality(t *testing.T) {
	var svc service.CipherService

	plaintext := "1234"
	k1 := "0x00000000000000000000000000000000"
	k2 := "0x000102030405060708090a0b0c0d0e0f"
	n1 := "0x000000000000000000000000"
	n2 := "0x000102030405060708090a0b"

	enc := func(key, nonce string) string {
		resp, err := svc.Encrypt(&model.EncryptionRequest{
			Plaintext: plaintext, IsPlaintextHex: true,
			Nonce: nonce, Key: key,
		})
		if err != nil {
			t.Fatalf("Encrypt() failed: %v", err)
		}
		return resp.Ciphertext
	}

	c1 := enc(k1, n1)
	c2 := enc(k2, n1)
	c3 := enc(k1, n2)

	dec := func(key, nonceCiphertext string) string {
		resp, err := svc.Decrypt(&model.DecryptionRequest{
			NonceCiphertext: nonceCiphertext, Key: key,
		})
		if err != nil {
			t.Fatalf("Decrypt() failed: %v", err)
		}
		return resp.Plaintext
	}

	// a) Decrypt C with (K1,V1) -> P
	p1 := dec(k1, c1)
	t.Logf("2a) ciphertext=%s, key=%s, plaintext=%s", c1, k1, p1)
	if p1 != plaintext {
		t.Errorf("decrypt C with (K1,V1) failed: got %s, want %x", p1, plaintext)
	}

	// b) Decrypt C' with (K2,V1) -> P
	p2 := dec(k2, c2)
	t.Logf("2b) ciphertext=%s, key=%s, plaintext=%s", c2, k2, p2)
	if p2 != plaintext {
		t.Errorf("decrypt C' with (K2,V1) failed: got %s, want %x", p2, plaintext)
	}

	// c) Decrypt C'' with (K1,V2) -> P
	p3 := dec(k1, c3)
	t.Logf("2c) ciphertext=%s, key=%s, plaintext=%s", c3, k1, p3)
	if p3 != plaintext {
		t.Errorf("decrypt C'' with (K1,V2) failed: got %s, want %x", p3, plaintext)
	}
}

// TestService_DecryptionWrongInput implements Test Case 3:
// Decrypt C with wrong key and/or wrong IV; should get wrong plaintext.
func TestService_DecryptionWrongInput(t *testing.T) {
	var svc service.CipherService

	plaintext := "1234"
	k := "0x00000000000000000000000000000000"
	kPrime := "0x000102030405060708090a0b0c0d0e0f"
	n := "0x000000000000000000000000"
	nPrime := "0x000102030405060708090a0b"

	resp, err := svc.Encrypt(&model.EncryptionRequest{
		Plaintext: plaintext, IsPlaintextHex: true,
		Nonce: n, Key: k,
	})
	if err != nil {
		t.Fatalf("Encrypt() failed: %v", err)
	}
	c := resp.Ciphertext
	t.Logf("3) plaintext=%s, key=%s, nonce=%s, ciphertext=%s", plaintext, k, n, c)

	// a) Decrypt with wrong key K', correct IV V
	pWrongKey, err := svc.Decrypt(&model.DecryptionRequest{
		NonceCiphertext: c, Key: kPrime,
	})
	if err != nil {
		t.Fatalf("Decrypt with wrong key failed: %v", err)
	}
	t.Logf("3a) ciphertext=%s, wrongKey=%s, decrypted=%s", c, kPrime, pWrongKey.Plaintext)
	if pWrongKey.Plaintext == plaintext {
		t.Error("decrypt with wrong key should not recover original plaintext")
	}

	// b) Decrypt with correct key K, wrong IV V' (replace nonce portion)
	cBytes, _ := hex.DecodeString(c)
	wrongIVBytes := make([]byte, len(cBytes))
	copy(wrongIVBytes, cBytes)
	nPrimeBytes, _ := utils.Hex2Byte(nPrime)
	copy(wrongIVBytes[:12], nPrimeBytes)
	cWrongIV := hex.EncodeToString(wrongIVBytes)

	pWrongIV, err := svc.Decrypt(&model.DecryptionRequest{
		NonceCiphertext: cWrongIV, Key: k,
	})
	if err != nil {
		t.Fatalf("Decrypt with wrong IV failed: %v", err)
	}
	t.Logf("3b) ciphertext=%s, correctKey=%s, wrongNonce=%s, decrypted=%s", cWrongIV, k, nPrime, pWrongIV.Plaintext)
	if pWrongIV.Plaintext == plaintext {
		t.Error("decrypt with wrong IV should not recover original plaintext")
	}

	// c) Decrypt with wrong key K' and wrong IV V'
	pWrongBoth, err := svc.Decrypt(&model.DecryptionRequest{
		NonceCiphertext: cWrongIV, Key: kPrime,
	})
	if err != nil {
		t.Fatalf("Decrypt with wrong key and IV failed: %v", err)
	}
	t.Logf("3c) ciphertext=%s, wrongKey=%s, wrongNonce=%s, decrypted=%s", cWrongIV, kPrime, nPrime, pWrongBoth.Plaintext)
	if pWrongBoth.Plaintext == plaintext {
		t.Error("decrypt with wrong key and wrong IV should not recover original plaintext")
	}
}

// TestService_ErrorPropagation implements Test Case 5:
// Bit-flip, bit-insertion, and bit-deletion errors in ciphertext.
func TestService_ErrorPropagation(t *testing.T) {
	var svc service.CipherService

	plaintext := "1234"
	k := "000102030405060708090a0b0c0d0e0f"
	n := "000102030405060708090a0b"

	resp, err := svc.Encrypt(&model.EncryptionRequest{
		Plaintext: plaintext, IsPlaintextHex: true,
		Nonce: n, Key: k,
	})
	if err != nil {
		t.Fatalf("Encrypt() failed: %v", err)
	}

	// Decode the hex ciphertext (nonce || actual ciphertext).
	// Use three-index slice to prevent append from aliasing into ct.
	full, _ := hex.DecodeString(resp.Ciphertext)
	nonce := full[:12:12]
	ct := make([]byte, len(full)-12)
	copy(ct, full[12:])

	// Helper: decrypt modified bytes.
	decrypt := func(modified []byte) string {
		combined := make([]byte, 0, len(nonce)+len(modified))
		combined = append(combined, nonce...)
		combined = append(combined, modified...)
		dresp, derr := svc.Decrypt(&model.DecryptionRequest{
			NonceCiphertext: hex.EncodeToString(combined), Key: k,
		})
		if derr != nil {
			t.Fatalf("Decrypt() failed: %v", derr)
		}
		return dresp.Plaintext
	}

	t.Run("bit-flip", func(t *testing.T) {
		ct2 := make([]byte, len(ct))
		copy(ct2, ct)
		ct2[0] ^= 0x08

		t.Logf("5a) original ciphertext=%s", resp.Ciphertext)
		combined := make([]byte, 0, len(nonce)+len(ct2))
		combined = append(combined, nonce...)
		combined = append(combined, ct2...)
		t.Logf("5a) flipped hex=%s", hex.EncodeToString(combined))

		p := decrypt(ct2)
		t.Logf("5a) decrypted plaintext=%s", p)
		if p == plaintext {
			t.Error("decryption of bit-flipped ciphertext should not match original plaintext")
		}

		// Verify only one byte differs.
		origPt, _ := hex.DecodeString(plaintext)
		gotPt, _ := hex.DecodeString(p)
		diffCount := 0
		for i := range origPt {
			if origPt[i] != gotPt[i] {
				diffCount++
			}
		}
		t.Logf("5a) differing byte count=%d (expected 1)", diffCount)
		if diffCount != 1 {
			t.Errorf("expected exactly 1 byte to differ, got %d", diffCount)
		}
	})

	t.Run("bit-insertion", func(t *testing.T) {
		bits := utils.Byte2Bits(ct)
		pos := 9
		bits = append(bits[:pos], append([]int{0}, bits[pos:]...)...)
		bits = bits[:len(bits)-1]
		ct2, err := utils.Bits2Byte(bits)
		if err != nil {
			t.Fatalf("Bits2Byte failed: %v", err)
		}

		fullHex := hex.EncodeToString(append(append([]byte{}, nonce...), ct...))
		modHex := hex.EncodeToString(append(append([]byte{}, nonce...), ct2...))
		t.Logf("5b) original ciphertext=%s", fullHex)
		t.Logf("5b) inserted bit 0 at pos %d: %s", pos, modHex)

		p := decrypt(ct2)
		t.Logf("5b) decrypted plaintext=%s", p)
		if p == plaintext {
			t.Error("decryption of bit-inserted ciphertext should not match original plaintext")
		}
	})

	t.Run("bit-deletion", func(t *testing.T) {
		bits := utils.Byte2Bits(ct)
		pos := 9
		bits = append(bits[:pos], bits[pos+1:]...)
		bits = append(bits, 0)
		ct2, err := utils.Bits2Byte(bits)
		if err != nil {
			t.Fatalf("Bits2Byte failed: %v", err)
		}

		fullHex := hex.EncodeToString(append(append([]byte{}, nonce...), ct...))
		modHex := hex.EncodeToString(append(append([]byte{}, nonce...), ct2...))
		t.Logf("5c) original ciphertext=%s", fullHex)
		t.Logf("5c) deleted bit at pos %d: %s", pos, modHex)

		p := decrypt(ct2)
		t.Logf("5c) decrypted plaintext=%s", p)
		if p == plaintext {
			t.Error("decryption of bit-deleted ciphertext should not match original plaintext")
		}
	})
}
