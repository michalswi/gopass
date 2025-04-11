package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

const secretLength = 32

// Generate256BitSecret generates a cryptographically secure 256-bit (32-byte) random secret.
func Generate256BitSecret() ([]byte, error) {
	secret := make([]byte, secretLength)
	_, err := io.ReadFull(rand.Reader, secret)
	// _, err := rand.Read(secret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random secret: %w", err)
	}
	return secret, nil
}
