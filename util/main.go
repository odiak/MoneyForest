package util

import (
	"crypto/rand"

	uuid "github.com/satori/go.uuid"
)

func StringPtr(s string) *string {
	return &s
}

func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

func RandomStr(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
