package service

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/grain"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

type EncryptionService struct{}

func newEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (e *EncryptionService) Encrypt(req *model.EncryptionRequest) (*model.EncryptionResp, error) {
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
