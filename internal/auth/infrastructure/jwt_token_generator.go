package authinfra

import (
	authdom "octodome/internal/auth/domain"
	userdom "octodome/internal/user/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtTokenGenerator struct {
}

func NewJwtTokenGenerator() authdom.AuthTokenGenerator {
	return &jwtTokenGenerator{}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (j *jwtTokenGenerator) GenerateToken(user *userdom.User) (string, error) {
	claims := authdom.UserClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "octodome",
			Subject:   "",
			Audience:  jwt.ClaimStrings{},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(24 * time.Hour)},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
