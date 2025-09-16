package repository

import (
	"octodome/internal/auth/domain"
	modelinfra "octodome/internal/auth/internal/infrastructure/model"
	"octodome/internal/core/collection"

	"gorm.io/gorm"
)

type pgRole struct {
	db *gorm.DB
}

func NewPgRole(db *gorm.DB) *pgRole {
	return &pgRole{db: db}
}

func (r *pgRole) GetRolesByUserID(userID uint) ([]domain.RoleDTO, error) {
	var roles []modelinfra.Role

	result := r.db.Table("roles").
		Select("roles.name").
		Joins("join user_roles on user_roles.role = roles.name").
		Where("user_roles.user_id = ?", userID).
		Scan(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return collection.Map(roles, func(r modelinfra.Role) domain.RoleDTO {
		return domain.RoleDTO{
			Name: domain.RoleName(r.Name),
		}
	}), nil
}
