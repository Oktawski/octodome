package security

import (
	"os"
	"time"

	authdomshared "octodome.com/api/internal/auth/domain"
	authdom "octodome.com/api/internal/auth/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type jwtTokenGenerator struct {
}

func NewJwtTokenGenerator() authdom.AuthTokenGenerator {
	return &jwtTokenGenerator{}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (j *jwtTokenGenerator) GenerateToken(user *authdomshared.UserAuthDTO) (string, error) {
	claims := authdomshared.UserClaims{
		UserID: user.ID,
		Email:  user.Email,
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
