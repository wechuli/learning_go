package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func CryptoRandPlay() {

	fmt.Println(generatePass(20))
}

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.RawURLEncoding.EncodeToString(b), err
}
