package keys

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/grain"
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
	// Wrapped stores the key after KDF and AES-128-CCM encryption.
	Wrapped    []byte
	Salt       []byte
	Iterations int
	Nonce      []byte
	Tag        []byte
}

// hexKey stores `Keys` information in hex string format.
type hexKey struct {
	Salt       string `json:"salt"`
	Iterations int    `json:"iteration"`
	Nonce      string `json:"nonce"`
	Tag        string `json:"tag"`
	Wrapped    string `json:"wrapped"`
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
	// Check the length of `kGrain`.
	if len(kGrain) != grain.KeyGrainBitLength/8 {
		return fmt.Errorf("key length should be %v byte, got: %v byte", grain.KeyGrainBitLength/8, len(kGrain))
	}

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

// Encode transforms bytes to hex, marshals the key to JSON format and encodes using base64.
func (k *Keys) Encode() (string, error) {
	hexK := keys2HexKey(k)

	// Marshal JSON.
	b, err := json.Marshal(hexK)
	if err != nil {
		return "", err
	}

	// base64 encode.
	encoded := base64.StdEncoding.EncodeToString(b)
	return encoded, nil
}

// Decode unmarshals json (byte format), and update the `Keys` object.
func (k *Keys) Decode(b64 string) error {
	jsonByte, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return fmt.Errorf("invalid file content: %w", err)
	}

	var hk hexKey
	if err := json.Unmarshal(jsonByte, &hk); err != nil {
		return fmt.Errorf("invalid json format: %w", err)
	}

	result, err := hexKey2Keys(&hk)
	if err != nil {
		return fmt.Errorf("invalid hex: %w", err)
	}

	*k = *result
	return nil
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

func keys2HexKey(k *Keys) *hexKey {
	return &hexKey{
		Iterations: k.Iterations,
		Salt:       hex.EncodeToString(k.Salt),
		Nonce:      hex.EncodeToString(k.Nonce),
		Tag:        hex.EncodeToString(k.Tag),
		Wrapped:    hex.EncodeToString(k.Wrapped),
	}
}

func hexKey2Keys(hk *hexKey) (*Keys, error) {
	salt, err := utils.Hex2Byte(hk.Salt)
	if err != nil {
		return nil, fmt.Errorf("invalid hex salt: %w", err)
	}

	nonce, err := utils.Hex2Byte(hk.Nonce)
	if err != nil {
		return nil, fmt.Errorf("invalid hex nonce: %w", err)
	}

	tag, err := utils.Hex2Byte(hk.Tag)
	if err != nil {
		return nil, fmt.Errorf("invalid hex tag: %w", err)
	}

	wrapped, err := utils.Hex2Byte(hk.Wrapped)
	if err != nil {
		return nil, fmt.Errorf("invalid hex wrapped key: %w", err)
	}

	return &Keys{
		Iterations: hk.Iterations,
		Salt:       salt,
		Nonce:      nonce,
		Tag:        tag,
		Wrapped:    wrapped,
	}, nil
}
