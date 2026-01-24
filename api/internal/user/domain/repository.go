package domain

import (
	"context"

	"octodome.com/shared/valuetype"
)

type Repository interface {
	GetByID(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) (uint, error)
	ExistsByEmail(ctx context.Context, email valuetype.Email) (bool, error)
}
