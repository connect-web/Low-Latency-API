package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateNonce creates a secure random nonce
func GenerateNonce() (string, error) {
	b := make([]byte, 16) // 16 bytes will generate a 24-character nonce when base64 encoded
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
