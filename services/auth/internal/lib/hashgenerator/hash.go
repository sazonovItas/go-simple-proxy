package hashgenerator

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func NewHash() (string, error) {
	const op = "lib.hashgenerator.NewHash"

	b := make([]byte, 60)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("%s: failed to generate random bytes: %w", op, err)
	}

	hash := sha256.New()
	_, err = hash.Write(b)
	if err != nil {
		return "", fmt.Errorf("%s: failed to compute hash: %w", op, err)
	}

	return string(hash.Sum(nil)), nil
}
