package hasher

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// Generate some hash or token with timestamp
func NewRandomHash() (string, error) {
	const op = "lib.hashgenerator.NewHash"

	hash := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
	return hash, nil
}