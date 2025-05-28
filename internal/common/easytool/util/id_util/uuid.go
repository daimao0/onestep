package id_util

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

// charset random charset to generate uuid
var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateShortUUID use to random index for charset to generate uuid
func GenerateShortUUID(length int) (string, error) {
	//
	if length <= 0 {
		return "", errors.New("length must be positive")
	}
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(bytes), nil
}

// GenerateUUID use to generate uuid by google uuid
func GenerateUUID() string {
	return uuid.NewString()
}

// GenerateUUIDWithoutSeparator generate uuid without separator
func GenerateUUIDWithoutSeparator() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
