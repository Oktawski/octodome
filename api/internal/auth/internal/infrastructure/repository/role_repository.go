package repository

import (
	"fmt"
	"octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/auth/internal/infrastructure/model"
	"octodome.com/shared/collection"

	"gorm.io/gorm"
)

type pgRole struct {
	db *gorm.DB
}

func NewPgRole(db *gorm.DB) *pgRole {
	return &pgRole{db: db}
}

func (r *pgRole) GetRolesByUserID(userID uint) ([]domain.RoleDTO, error) {
	var roles []model.Role

	result := r.db.Table("roles").
		Select("roles.name").
		Joins("join user_roles on user_roles.role_id = roles.name").
		Where("user_roles.user_id = ?", userID).
		Scan(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return collection.Map(roles, func(r model.Role) domain.RoleDTO {
		return domain.RoleDTO{
			Name: domain.RoleName(r.Name),
		}
	}), nil
}

func (r *pgRole) AssignRole(role domain.RoleName, userID uint) error {
	var dbRole model.Role
	if err := r.db.Model(&model.Role{}).
		Where("name = ?", role).
		First(&dbRole).
		Error; err != nil {
		return fmt.Errorf("role %s does not exist", role)
	}

	var count int64
	if err := r.db.Model(&model.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, role).
		Count(&count).
		Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("user already has role %s", role)
	}

	userRole := &model.UserRole{
		UserID: userID,
		RoleID: string(role),
	}

	if err := r.db.Create(&userRole).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgRole) UnassignRole(role domain.RoleName, userID uint) error {
	err := r.db.
		Model(&model.UserRole{}).
		Where("role_id = ? AND user_id = ?", role, userID).
		Delete(nil).
		Error
	return err
}

func (r *pgRole) SyncRoles(
	rolesToAdd []domain.RoleName,
	rolesToRemove []domain.RoleName,
	userID uint,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if len(rolesToAdd) > 0 {
			userRoles := make([]model.UserRole, 0, len(rolesToAdd))
			for _, role := range rolesToAdd {
				userRoles = append(userRoles, model.UserRole{
					UserID: userID,
					RoleID: string(role),
				})
			}
			if err := tx.Create(&userRoles).Error; err != nil {
				return err
			}
		}

		if len(rolesToRemove) > 0 {
			if err := tx.
				Where("user_id = ? AND role_id IN ?", userID, rolesToRemove).
				Delete(&model.UserRole{}).
				Error; err != nil {
				return err
			}
		}

		return nil
	})
}
