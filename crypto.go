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

// AES encryption with HMAC
func EncryptAES(key, message []byte) ([]byte, error) {
	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create the byte slice that will hold encrypted message
	cipherText := make([]byte, aes.BlockSize+len(message))

	// Generate the Initialization Vector nonce which is
	// stored at the beginning of the byte slice.
	// The IV is the same length as the AES blocksize
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Create the AES cipher stream
	cfb := cipher.NewCFBEncrypter(block, iv)

	// Generate the encrypted message and store it in
	// the remaining bytes after the IV nonce
	cfb.XORKeyStream(cipherText[aes.BlockSize:], message)

	// Create a new HMAC
	h := hmac.New(sha256.New, key)

	// Write the ciphertext to the HMAC
	h.Write(cipherText)

	// Append the HMAC to the ciphertext
	cipherText = h.Sum(cipherText)

	return cipherText, nil
}

// AES decryption with HMAC
func DecryptAES(key, cipherText []byte) ([]byte, error) {
	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Separate the HMAC from the ciphertext
	hmacSize := sha256.Size
	if len(cipherText) < hmacSize {
		return nil, errors.New("ciphertext too short")
	}
	msgHMAC := cipherText[len(cipherText)-hmacSize:]
	cipherText = cipherText[:len(cipherText)-hmacSize]

	// Verify the HMAC
	h := hmac.New(sha256.New, key)
	h.Write(cipherText)
	expectedHMAC := h.Sum(nil)
	if !hmac.Equal(msgHMAC, expectedHMAC) {
		return nil, errors.New("HMAC verification failed")
	}

	// Separate the IV nonce from the encrypted message bytes
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message using the CFB block mode
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
