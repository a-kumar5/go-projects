package crypto

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// DeriveKeyFromPassphrase derives a 32-byte key using scrypt.
// Parameters: passphrase string and a 16+ byte salt.
// N (cpu/mem cost) = 2^15, r=8, p=1. Adjust for your environment.
func DeriveKeyFromPassphrase(passphrase string, salt []byte) ([]byte, error) {
	if len(salt) < 8 {
		return nil, fmt.Errorf("salt must be at least 8 bytes, got %d", len(salt))
	}
	N := 1 << 15 // 32768
	r := 8
	p := 1
	keyLen := 32
	return scrypt.Key([]byte(passphrase), salt, N, r, p, keyLen)
}

// ConstantTimeEqual helper.
func ConstantTimeEqual(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}
