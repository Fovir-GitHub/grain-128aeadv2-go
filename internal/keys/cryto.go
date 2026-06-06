package keys

import (
	"crypto/pbkdf2"
	"crypto/sha256"
)

// kdf uses PBKDF2-HMAC-SHA256 to derive a wrapping key.
func kdf(passphrase string, salt []byte, iter, keysize int) ([]byte, error) {
	return pbkdf2.Key(sha256.New, passphrase, salt, iter, keysize)
}
