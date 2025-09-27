package repository

import "octodome/internal/auth/domain"

type UserReader interface {
	GetUserAuthDTO(username string) (*domain.UserAuthDTO, error)
}
