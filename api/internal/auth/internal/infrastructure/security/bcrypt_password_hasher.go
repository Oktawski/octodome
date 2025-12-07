package security

import (
	"octodome.com/api/internal/auth/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() domain.PasswordHasher {
	return &bcryptPasswordHasher{}
}

func (h *bcryptPasswordHasher) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (h *bcryptPasswordHasher) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password))
}
