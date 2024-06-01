package entity

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	ID    string
	Email string
	Login string
	Role  string
}

type TokenClaims struct {
	*jwt.RegisteredClaims
	Info UserInfo
}
