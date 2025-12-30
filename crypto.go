package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

const (
	KeySize   = 32 // AES-256 key size
	NonceSize = 12 // Standard GCM nonce size to keep output stable
)

// EncryptAES encrypts a message using AES-256-GCM AEAD (authenticated encryption).
// Output format: [nonce || ciphertext || tag]
func EncryptAES(key, message []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, fmt.Errorf("invalid key length: must be %d bytes", KeySize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCMWithNonceSize(block, NonceSize)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	sealed := aead.Seal(make([]byte, 0, len(message)+aead.Overhead()), nonce, message, nil)

	out := make([]byte, 0, NonceSize+len(sealed))
	out = append(out, nonce...)
	out = append(out, sealed...)
	return out, nil
}

// DecryptAES decrypts data produced by EncryptAES and verifies its authenticity.
func DecryptAES(key, cipherText []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, fmt.Errorf("invalid key length: must be %d bytes", KeySize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCMWithNonceSize(block, NonceSize)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < NonceSize+aead.Overhead() {
		return nil, errors.New("ciphertext too short: missing nonce or tag")
	}

	nonce := cipherText[:NonceSize]
	sealed := cipherText[NonceSize:]

	return aead.Open(nil, nonce, sealed, nil)
}
