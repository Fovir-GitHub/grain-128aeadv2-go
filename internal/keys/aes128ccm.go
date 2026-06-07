package keys

import (
	"crypto/aes"
	"fmt"

	"github.com/pion/dtls/v2/pkg/crypto/ccm"
)

type AES128CCM struct {
	Nonce []byte
	AD    []byte
}

func NewAES128CCM(nonce, ad []byte) (*AES128CCM, error) {
	if len(nonce) != NonceSize {
		return nil, fmt.Errorf("nonce length in AES-128-CCM should be %v, got %v", NonceSize, len(nonce))
	}

	return &AES128CCM{
		Nonce: nonce,
		AD:    ad,
	}, nil
}

// Encrypt calculates ciphertext and tag using AES-128-CCM.
func (a *AES128CCM) Encrypt(key, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	c, err := ccm.NewCCM(block, TagSize, NonceSize)
	if err != nil {
		return nil, nil, err
	}

	output := c.Seal(nil, a.Nonce, plaintext, a.AD)
	ciphertext := output[:len(plaintext)]
	tag := output[len(plaintext):]

	return ciphertext, tag, nil
}

// Auth authenticates the input key with the `.key` file.
func (a *AES128CCM) Auth(key, ciphertext, tag []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	c, err := ccm.NewCCM(block, TagSize, NonceSize)
	if err != nil {
		return nil, err
	}

	combined := append(ciphertext, tag...)
	plaintext, err := c.Open(nil, a.Nonce, combined, a.AD)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
