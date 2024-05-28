package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user *entity.User, secret string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["login"] = user.Login
	claims["role"] = user.UserRole
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
