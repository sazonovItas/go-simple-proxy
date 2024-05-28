package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	authsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/auth"
)

const DefaultCost = bcrypt.DefaultCost

// hasher is implementing Hasher interface
type hasher struct {
	cost int
}

var _ authsvc.Hasher = (*hasher)(nil)

func New(cost int) *hasher {
	return &hasher{cost: cost}
}

// Password is implementing Hasher interface
func (h *hasher) PasswordHash(password string) ([]byte, error) {
	const op = "lib.hasher.PasswordHash"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return nil, fmt.Errorf("%s: failed generate hash %w", op, err)
	}

	return hashedPassword, nil
}

// Compare is implementing Hasher interface
func (h *hasher) Compare(hashedPassword string, password string) error {
	const op = "lib.hasher.Compare"

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
