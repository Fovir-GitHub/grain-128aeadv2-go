package model

type EncryptionRequest struct {
	Plaintext      string `json:"input"`
	IsPlaintextHex bool   `json:"isInputHex"`
	Nonce          string `json:"nonce"`
	Key            string `json:"key"`
}
