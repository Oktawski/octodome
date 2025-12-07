package domain

import (
	"octodome.com/api/internal/auth/domain"
)

type AuthTokenGenerator interface {
	GenerateToken(user *domain.UserAuthDTO) (string, error)
}
