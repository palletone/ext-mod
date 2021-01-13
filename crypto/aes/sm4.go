package aes

import (
	"crypto/cipher"
	"github.com/palletone/ext-mod/crypto/internal/sm4"
)
const BlockSize = sm4.BlockSize
func NewCipher(key []byte) (cipher.Block, error) {
	return sm4.NewCipher(key)
}

