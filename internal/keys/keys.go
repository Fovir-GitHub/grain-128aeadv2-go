package keys

import (
	"encoding/json"
	"fmt"
	"io"
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

// Wrap wraps `kGrain` with a given `passphrase` and associated data (`ad`).
//
// It generates a pair of random salt and nonce before wrapping the key, and calculate the corresponding `Tag` derived from AES-128-CCM algorithm.
func (k *Keys) Wrap(passphrase string, kGrain, ad []byte) error {
	// Generate random salt and nonce.
	salt, err := utils.RandomBytes(SaltSize)
	if err != nil {
		return fmt.Errorf("random salt generation failed: %w", err)
	}

	nonce, err := utils.RandomBytes(NonceSize)
	if err != nil {
		return fmt.Errorf("random nonce generation failed: %w", err)
	}

	// Call KDF to wrap `passphrase` and get  $K_{wrap}$.
	kWrap, err := kdf(passphrase, salt, k.Iterations, KeySize)
	if err != nil {
		return fmt.Errorf("kWrap generation failed: %w", err)
	}

	// Encrypt `kGrain` with `kWrap` using AES-128-CCM.
	a, err := NewAES128CCM(nonce, ad)
	if err != nil {
		return err
	}

	ciphertext, tag, err := a.Encrypt(kWrap, kGrain)
	if err != nil {
		return fmt.Errorf("encrypt kGrain using kWrap failed: %w", err)
	}

	// Store values.
	k.Wrapped = ciphertext
	k.Nonce = nonce
	k.Tag = tag
	k.Salt = salt

	return nil
}

// Unwrap recovers the wrapped key with given `passphrase`,
// and authenticates the integrity using the given associated data (`ad`).
func (k *Keys) Unwrap(passphrase string, ad []byte) (kGrain []byte, err error) {
	// Derive kWrap.
	kWrap, err := kdf(passphrase, k.Salt, k.Iterations, KeySize)
	if err != nil {
		return nil, err
	}

	a, err := NewAES128CCM(k.Nonce, ad)
	if err != nil {
		return nil, err
	}

	return a.Auth(kWrap, k.Wrapped, k.Tag)
}

func (k *Keys) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(k)
}

func (k *Keys) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close() //nolint
	return json.NewEncoder(f).Encode(k)
}

func (k *Keys) ReadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close() //nolint
	return json.NewDecoder(f).Decode(k)
}
