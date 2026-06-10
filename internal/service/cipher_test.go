package service_test

import (
	"testing"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/service"
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
