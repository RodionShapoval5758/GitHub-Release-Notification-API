package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateToken(lengths int) (string, error) {
	bytes := make([]byte, lengths)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("generating random bytes: %w", err)
	}

	token := base64.RawURLEncoding.EncodeToString(bytes)

	return token, nil
}
