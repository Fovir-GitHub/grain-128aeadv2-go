package grain

import (
	"fmt"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
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

	if len(k) != 128 {
		return nil, fmt.Errorf("key length should be 128 bits (len=%v)", len(key))
	}

	if len(n) != 96 {
		return nil, fmt.Errorf("key length should be 128 bits (len=%v)", len(nonce))
	}

	return &Grain128AEADV2{
		key:   k,
		nonce: n,
		LFSR:  make([]int, 0, 128),
		NFSR:  make([]int, 128),
	}, nil
}
