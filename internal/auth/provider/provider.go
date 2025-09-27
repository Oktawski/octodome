package provider

import (
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/infrastructure/repository"

	"gorm.io/gorm"
)

func ProvideRoleReader(db *gorm.DB) domain.RoleRepository {
	return repository.NewPgRole(db)
}
