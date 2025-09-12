package domain

import userdom "octodome/internal/user/domain"

type Repository interface {
	GetUserByUsername(username string) (*userdom.User, error)
}
