package service

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/keys"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

type KeyManagementService struct{}

func newKeyManagement() *KeyManagementService {
	return &KeyManagementService{}
}

func (s *KeyManagementService) WrapKey(req *model.WrapKeyRequest) (*keys.Keys, error) {
	// Parse kGrain and AD.
	var adBytes []byte
	var kGrainBytes []byte
	var err error

	kGrainBytes, err = utils.Hex2Byte(req.Key)
	if err != nil {
		return nil, fmt.Errorf("invalid hex key: %v", req.Key)
	}

	if req.IsHex {
		adBytes, err = utils.Hex2Byte(req.AD)
		if err != nil {
			return nil, fmt.Errorf("invalid hex AD: %v", req.AD)
		}
	} else {
		adBytes = []byte(req.AD)
	}

	k := keys.New()
	if err = k.Wrap(req.Password, kGrainBytes, adBytes); err != nil {
		return nil, fmt.Errorf("wrap key failed: %w", err)
	}

	return k, nil
}

func (s *KeyManagementService) UnwrapKey(req *model.UnwrapKeyRequest) (*model.UnwrapKeyResp, error) {
	// Decode base64 content.
	jsonByte, err := base64.StdEncoding.DecodeString(req.Base64Content)
	if err != nil {
		return nil, fmt.Errorf("invalid file content: %w", err)
	}

	// Unmarshal the json into `keys.Keys`.
	var k keys.Keys
	if err := json.Unmarshal(jsonByte, &k); err != nil {
		return nil, fmt.Errorf("invalid json format: %w", err)
	}

	// Unwrap `k` to obtain `kGrain` and auth the data.
	kGrain, err := k.Unwrap(req.Passphrase, []byte(req.AD))
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// Transform kGrain into hex format.
	key := hex.EncodeToString(kGrain)

	return &model.UnwrapKeyResp{
		Key: key,
	}, nil
}
