package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateBase64(key string) string {
	bytes := []byte(key)
	hash := base64.StdEncoding.EncodeToString(bytes)
	return hash
}

func GenerateSHA256(key string) string {
	bytes := []byte(key)
	hash := sha256.Sum256(bytes)
	val := fmt.Sprintf("%x", hash)
	return val
}