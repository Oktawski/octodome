package repository

import "octodome.com/api/internal/auth/domain"

type UserReader interface {
	GetUserAuthDTO(username string) (*domain.UserAuthDTO, error)
}
