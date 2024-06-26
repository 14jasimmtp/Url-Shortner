package utils

import (
	"crypto/rand"
	"math/big"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const length = 5

// Function to generate a random string of specified length
func GenerateRandomString() (string, error) {
	b := make([]byte, 5)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[num.Int64()]
	}
	return string(b), nil
}
