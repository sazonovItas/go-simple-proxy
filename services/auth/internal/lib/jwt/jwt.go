package jwt

import (
	"errors"
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
func NewToken(
	userInfo entity.UserInfo,
	secret string,
	duration time.Duration,
) (string, error) {
	const op = "lib.jwt.NewToken"

	token := jwt.New(jwt.SigningMethodHS256)

	exp := time.Now().Add(duration)
	tokenClaims := &entity.TokenClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   userInfo.ID,
		},
		Info: userInfo,
	}

	token.Claims = tokenClaims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenString, nil
}

// ValidateToken validates token and returns token claims
func ValidateToken(tokenString, secret string) (entity.UserInfo, error) {
	const op = "lib.jwt.ValidateToken"

	token, err := jwt.NewParser(jwt.WithExpirationRequired()).
		ParseWithClaims(tokenString, &entity.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return entity.UserInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	if claims, ok := token.Claims.(*entity.TokenClaims); ok {
		return claims.Info, nil
	}

	return entity.UserInfo{}, fmt.Errorf(
		"%s: %w",
		op,
		errors.New("failed to get claims from token"),
	)
}
