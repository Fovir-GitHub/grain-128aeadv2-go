package service

import (
	"encoding/hex"
	"fmt"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/keys"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

type KeyManagementService struct{}

func newKeyManagement() *KeyManagementService {
	return &KeyManagementService{}
}

func (s *KeyManagementService) WrapKey(req *model.WrapKeyRequest) (*model.WrapKeyResp, error) {
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

	encoded, err := k.Encode()
	if err != nil {
		return nil, fmt.Errorf("encode key failed: %w", err)
	}

	return &model.WrapKeyResp{
		Key: encoded,
	}, nil
}

func (s *KeyManagementService) UnwrapKey(req *model.UnwrapKeyRequest) (*model.UnwrapKeyResp, error) {
	// Decode base64 json.
	k := keys.New()
	if err := k.Decode(req.Base64Content); err != nil {
		return nil, err
	}

	// Transform `req.AD` into bytes according to the format.
	var ad []byte
	var err error
	if req.IsHex {
		if ad, err = utils.Hex2Byte(req.AD); err != nil {
			return nil, fmt.Errorf("invalid hex ad: %w", err)
		}
	} else {
		ad = []byte(req.AD)
	}

	// Unwrap `k` to obtain `kGrain` and auth the data.
	kGrain, err := k.Unwrap(req.Passphrase, ad)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// Transform kGrain into hex format.
	key := hex.EncodeToString(kGrain)

	return &model.UnwrapKeyResp{
		Key: key,
	}, nil
}
