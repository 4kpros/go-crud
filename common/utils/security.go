package utils

import (
	"golang.org/x/crypto/sha3"
)

// Encrypt using SHA3-512
func EncryptValue(value string) string {
	h := sha3.New512()
	h.Write([]byte(value))
	sum := h.Sum(nil)
	return string(sum)
}

func DecryptValue(value string) string {
	h := sha3.New512()
	h.Write([]byte(value))
	sum := h.Sum(nil)
	return string(sum)
}
