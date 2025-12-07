package domain

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}
