package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"octodome.com/shared/valuetype"
)

type UserClaims struct {
	UserID uint
	Email  valuetype.Email
	jwt.RegisteredClaims
}
