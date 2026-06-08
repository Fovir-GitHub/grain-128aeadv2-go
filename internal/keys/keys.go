package keys

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"
)

const (
	SaltSize   = 16
	NonceSize  = 12
	TagSize    = 16
	Iterations = 600_000
	KeySize    = 16
)

type Keys struct {
	Salt       []byte `json:"salt"`
	Iterations int    `json:"iteration"`
	Nonce      []byte `json:"nonce"`
	Tag        []byte `json:"tag"`

	// Wrapped stores the key after KDF and AES-128-CCM encryption.
	Wrapped []byte `json:"wrapped"`
}

func New() *Keys {
	return &Keys{
		Iterations: Iterations,
	}
}
