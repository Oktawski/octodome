package domain

import "context"

type Repository interface {
	GetByID(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) (uint, error)
	ExistsByEmailOrUsername(ctx context.Context, email string, username string) (bool, error)
}
