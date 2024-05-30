package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

var (
	ErrTokenExpired   = jwt.ErrTokenExpired
	ErrTokenMalformed = jwt.ErrTokenMalformed
)

// NewToken creates new JWT token for given user and app.
func NewToken(user *entity.User, secret string, duration time.Duration) (string, error) {
	const op = "lib.jwt.NewToken"

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["login"] = user.Login
	claims["role"] = user.UserRole
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString, secret string) error {
	const op = "lib.jwt.ValidateToken"

	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
