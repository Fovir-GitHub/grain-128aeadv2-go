package grain

import (
	"fmt"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

const (
	KeyGrainBitLength         = 128
	FeedbackRegisterBitLength = 128
	NonceBitLength            = 96
)

type Grain128AEADV2 struct {
	key   []int // 128-bit
	nonce []int // 96-bit
	LFSR  []int // Aka S_t
	NFSR  []int // Aka B_t
}

func New(key, nonce []byte) (*Grain128AEADV2, error) {
	k := utils.Byte2Bits(key)
	n := utils.Byte2Bits(nonce)

	if len(k) != KeyGrainBitLength {
		return nil, fmt.Errorf("key length should be %v bits (len=%v)", KeyGrainBitLength, len(k))
	}

	if len(n) != 96 {
		return nil, fmt.Errorf("nonce length should be %v bits (len=%v)", NonceBitLength, len(n))
	}

	return &Grain128AEADV2{
		key:   k,
		nonce: n,
		LFSR:  make([]int, 0, FeedbackRegisterBitLength),
		NFSR:  make([]int, FeedbackRegisterBitLength),
	}, nil
}
