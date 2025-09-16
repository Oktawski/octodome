package domain

import (
	"octodome/internal/auth/domain"
)

type AuthTokenGenerator interface {
	GenerateToken(user *domain.UserAuthDTO) (string, error)
}
