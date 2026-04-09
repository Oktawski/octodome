package repository

import (
	"context"

	"octodome.com/api/internal/auth/domain"
	"octodome.com/shared/valuetype"
)

type UserReader interface {
	GetUserAuthDTO(ctx context.Context, email string) (*domain.UserAuthDTO, error)
	ExistsByEmailOrUsername(ctx context.Context, email valuetype.Email, username string) (bool, error)
}
