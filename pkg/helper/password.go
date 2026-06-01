package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordHash(password, hash string) string {
	sum := sha256.Sum256([]byte(password + hash))
	return hex.EncodeToString(sum[:])
}
