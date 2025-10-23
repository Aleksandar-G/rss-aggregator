package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Generates a cryptographically secure random hexadecimal string of a given byte length.
func GenerateRandomHexString(byteLength int) (string, error) {
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		// return "", fmt.Errorf("failed to read random bytes: %w", err)
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Computes the SHA256 hash of a given string and returns it as a hexadecimal string.
func HashSHA256String(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateApiKey() (string, error) {
	randomBytesLength := 16
	apiKey, err := GenerateRandomHexString(randomBytesLength)
	if err != nil {
		return "", err
	}

	sha256ApiKey := HashSHA256String(apiKey)

	return sha256ApiKey, nil
}
