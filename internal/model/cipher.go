package model

type EncryptionRequest struct {
	Plaintext      string `json:"plaintext"`
	IsPlaintextHex bool   `json:"isInputHex"`
	Nonce          string `json:"nonce"`
	Key            string `json:"key"`
}

type EncryptionResp struct {
	// Ciphertext is `IV || ciphertext` in hex format.
	Ciphertext string `json:"output"`

	// States in hex format.

	LoadedLFSR string `json:"loadedLFSR"`
	LoadedNFSR string `json:"loadedNFSR"`
	InitLFSR   string `json:"initLFSR"`
	InitNFSR   string `json:"initNFSR"`
}

type DecryptionRequest struct {
	// NonceCiphertext is `IV || ciphertext` in hex format.
	NonceCiphertext string `json:"ciphertext"`
	Key             string `json:"key"`
}

type DecryptionResp struct {
	Plaintext string `json:"output"`
	Nonce     string `json:"nonce"`

	// States in hex format.

	LoadedLFSR string `json:"loadedLFSR"`
	LoadedNFSR string `json:"loadedNFSR"`
	InitLFSR   string `json:"initLFSR"`
	InitNFSR   string `json:"initNFSR"`
}
