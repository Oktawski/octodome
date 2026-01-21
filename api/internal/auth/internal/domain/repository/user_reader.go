package repository

import (
	"context"

	"octodome.com/api/internal/auth/domain"
)

type UserReader interface {
	GetUserAuthDTO(ctx context.Context, email string) (*domain.UserAuthDTO, error)
}
