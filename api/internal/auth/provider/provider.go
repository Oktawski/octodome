package provider

import (
	domain "octodome.com/api/internal/auth/internal/domain/repository"
	"octodome.com/api/internal/auth/internal/infrastructure/repository"

	"gorm.io/gorm"
)

func ProvideRoleReader(db *gorm.DB) domain.RoleRepository {
	return repository.NewPgRole(db)
}
