package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken256() (token string, err error) {
	bytes := make([]byte, 128)

	if _, err = rand.Read(bytes); err != nil {
		return
	}

	token = hex.EncodeToString(bytes)

	return
}

func GenerateToken64() (token string, err error) {
	bytes := make([]byte, 32)

	if _, err = rand.Read(bytes); err != nil {
		return
	}

	token = hex.EncodeToString(bytes)

	return
}
