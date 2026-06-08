package service

import (
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
