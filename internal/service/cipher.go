package service

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/grain"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

type CipherService struct{}

func newCipherService() *CipherService {
	return &CipherService{}
}

// Encrypt implements the encryption service.
func (c *CipherService) Encrypt(req *model.EncryptionRequest) (*model.EncryptionResp, error) {
	// Trim space around key and nonce.
	req.Key = strings.TrimSpace(req.Key)
	req.Nonce = strings.TrimSpace(req.Nonce)

	// Transform plaintext, key, and nonce, from ASCII or hex string into bytes.
	var plaintext []byte
	var err error
	if req.IsPlaintextHex {
		plaintext, err = utils.Hex2Byte(req.Plaintext)
		if err != nil {
			return nil, fmt.Errorf("invalid hex plaintext: %v", req.Plaintext)
		}
	} else {
		plaintext = []byte(req.Plaintext)
	}

	kGrain, err := utils.Hex2Byte(req.Key)
	if err != nil {
		return nil, fmt.Errorf("invalid hex key: %v", req.Key)
	}

	nonce, err := utils.Hex2Byte(req.Nonce)
	if err != nil {
		return nil, fmt.Errorf("invalid hex nonce: %v", req.Nonce)
	}

	// Create `Grain128AEADV2` using key and nonce.
	g, err := grain.New(kGrain, nonce)
	if err != nil {
		return nil, fmt.Errorf("create Grain128AEADV2 failed: %w", err)
	}

	// Get states of LFSR and NFSR, and encrypt plaintext.
	loadedLFSR, loadedNFSR, initLFSR, initNFSR := g.Init()
	ciphertextByte := g.Encrypt(plaintext)

	// Prepend IV to `ciphertextByte` and encode to hex.
	ciphertextByte = append(nonce, ciphertextByte...)
	ciphertext := hex.EncodeToString(ciphertextByte)

	return &model.EncryptionResp{
		Ciphertext: ciphertext,
		LoadedLFSR: loadedLFSR,
		LoadedNFSR: loadedNFSR,
		InitLFSR:   initLFSR,
		InitNFSR:   initNFSR,
	}, nil
}

// Decrypt implements the decrytion service.
func (c *CipherService) Decrypt(req *model.DecryptionRequest) (*model.DecryptionResp, error) {
	const nonceByteLength = grain.NonceBitLength / 8

	// Trim spaces around key and ciphertext.
	req.NonceCiphertext = strings.TrimSpace(req.NonceCiphertext)
	req.Key = strings.TrimSpace(req.Key)

	// Transform ciphertext and key from hex to bytes.
	nonceCiphertext, err := utils.Hex2Byte(req.NonceCiphertext)
	if err != nil {
		return nil, fmt.Errorf("invalid hex ciphertext: %v", req.NonceCiphertext)
	}

	kGrain, err := utils.Hex2Byte(req.Key)
	if err != nil {
		return nil, fmt.Errorf("invalid hex key: %v", req.Key)
	}

	// Check ciphertext and key length.
	if len(nonceCiphertext) < nonceByteLength {
		return nil, fmt.Errorf("invalid ciphertext: too short")
	}

	// Extract nonce and ciphertext.
	nonce := nonceCiphertext[:nonceByteLength]
	ciphertext := nonceCiphertext[nonceByteLength:]

	// Create `Grain128AEADV2` using key and nonce.
	g, err := grain.New(kGrain, nonce)
	if err != nil {
		return nil, fmt.Errorf("create Grain128AEADV2 failed: %w", err)
	}

	// Get states of LFSR, NFSR, and decrypt ciphertext.
	loadedLFSR, loadedNFSR, initLFSR, initNFSR := g.Init()
	plaintextByte := g.Decrypt(ciphertext)

	// Encode plaintext to hex.
	plaintext := hex.EncodeToString(plaintextByte)

	// Get hex format nonce.
	nonceHexStr := hex.EncodeToString(nonce)

	return &model.DecryptionResp{
		Plaintext:  plaintext,
		Nonce:      nonceHexStr,
		LoadedLFSR: loadedLFSR,
		LoadedNFSR: loadedNFSR,
		InitLFSR:   initLFSR,
		InitNFSR:   initNFSR,
	}, nil
}
