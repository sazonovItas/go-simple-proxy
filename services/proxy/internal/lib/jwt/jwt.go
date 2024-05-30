package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func parseUnverified(tokenString string) (map[string]interface{}, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("cannot convert to map claims")
	}
}

func GetUserID(tokenString string) (string, error) {
	const op = "lib.jwt.GetUserID"

	claims, err := parseUnverified(tokenString)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if id, ok := claims["uid"].(string); ok {
		return id, nil
	} else {
		return "", errors.New("cannot parse uid from token string")
	}
}
