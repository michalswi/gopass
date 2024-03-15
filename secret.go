package main

import "crypto/rand"

func Generate256BitSecret() ([]byte, error) {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}
