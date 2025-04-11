package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

// EncryptAES encrypts a message using AES-CFB with a random IV and appends an HMAC-SHA256.
// The output format is: [IV (16 bytes) || ciphertext || HMAC (32 bytes)].
func EncryptAES(key, message []byte) ([]byte, error) {

	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate the Initialization Vector (IV),
	// which will be prepended to the ciphertext
	// in the final output.
	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Create the AES cipher stream
	cfb := cipher.NewCFBEncrypter(block, iv)
	// Create the byte slice that will hold encrypted message
	cipherText := make([]byte, len(message))
	cfb.XORKeyStream(cipherText, message)

	// Create a new HMAC
	h := hmac.New(sha256.New, key)
	// Write the IV nonce to the HMAC
	h.Write(iv)
	// Write the ciphertext to the HMAC
	h.Write(cipherText)
	// Append the HMAC to the IV and ciphertext to form
	// the final output: [IV || ciphertext || HMAC].
	mac := h.Sum(nil)

	return append(append(iv, cipherText...), mac...), nil
}

// DecryptAES decrypts a message encrypted by EncryptAES and verifies its HMAC.
// It expects the input format: [IV (16 bytes) || ciphertext || HMAC (32 bytes)].
func DecryptAES(key, cipherText []byte) ([]byte, error) {

	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	if len(cipherText) < aes.BlockSize+sha256.Size {
		return nil, errors.New("ciphertext too short: must include IV and HMAC")
	}

	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Separate the HMAC from the ciphertext
	hmacSize := sha256.Size
	msgHMAC := cipherText[len(cipherText)-hmacSize:]
	cipherText = cipherText[:len(cipherText)-hmacSize]

	// Verify the HMAC to ensure the integrity of
	// the IV and ciphertext.
	h := hmac.New(sha256.New, key)
	h.Write(cipherText) // Includes IV and ciphertext
	expectedHMAC := h.Sum(nil)
	if !hmac.Equal(msgHMAC, expectedHMAC) {
		return nil, errors.New("HMAC verification failed")
	}

	// Separate the IV from the ciphertext and use it
	// to initialize the decryption process.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message using the CFB block mode
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
