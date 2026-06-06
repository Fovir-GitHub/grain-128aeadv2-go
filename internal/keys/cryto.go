package keys

import (
	"crypto/pbkdf2"
	"crypto/sha256"
)

func kdf(passphrase string, salt []byte, iter, keysize int) ([]byte, error) {
	return pbkdf2.Key(sha256.New, passphrase, salt, iter, keysize)
}
