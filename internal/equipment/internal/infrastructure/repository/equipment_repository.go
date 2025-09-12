package repo

import (
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core"
	"octodome/internal/core/collection"
	domain "octodome/internal/equipment/internal/domain/equipment"
	"octodome/internal/equipment/internal/infrastructure/model"

	"gorm.io/gorm"
)

type pgEquipmentRepository struct {
	db *gorm.DB
}

func NewPgEquipmentRepository(db *gorm.DB) *pgEquipmentRepository {
	return &pgEquipmentRepository{db: db}
}

func (r *pgEquipmentRepository) GetList(
	page int,
	pageSize int,
	user authdom.UserContext,
) ([]domain.Equipment, int64, error) {

	var equipments []model.Equipment

	var count int64

	if dbError := r.db.
		Where("user_id = ?", user.ID).
		Scopes(core.Paginate(page, pageSize)).
		Find(&equipments).
		Count(&count).Error; dbError != nil {
		return nil, 0, dbError
	}

	equipmentList := collection.Map(
		equipments,
		func(e model.Equipment) domain.Equipment {
			return *e.ToDomain()
		},
	)

	return equipmentList, count, nil
}

func (r *pgEquipmentRepository) GetByID(
	id uint,
	user authdom.UserContext,
) (*domain.Equipment, error) {

	var eq *model.Equipment

	if dbError := r.db.
		Where("id = ? AND user_id = ?", id, user.ID).
		First(&eq).Error; dbError != nil {
		return nil, dbError
	}

	return eq.ToDomain(), nil
}

func (r *pgEquipmentRepository) Create(e *domain.Equipment) error {
	equipment := model.EquipmentFromDomain(e)

	if dbError := r.db.Create(&equipment).Error; dbError != nil {
		return dbError
	}

	return nil
}

func (r *pgEquipmentRepository) Update(e *domain.Equipment) error {
	if dbError := r.db.
		Model(&model.Equipment{}).
		Where("id = ? AND user_id = ?", e.ID, e.UserID).
		Updates(model.Equipment{
			Category:    e.Category,
			Description: e.Description,
			Name:        e.Name,
		}).Error; dbError != nil {
		return dbError
	}

	return nil
}

func (r *pgEquipmentRepository) Delete(id uint) error {
	if dbError := r.db.
		Model(&model.Equipment{}).
		Where("ID = ?", id).
		Delete(&model.Equipment{}).Error; dbError != nil {
		return dbError
	}
	return nil
}

func (r *pgEquipmentRepository) ExistsByNameAndType(
	name string,
	equipmentTypeID uint,
	userContext authdom.UserContext,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("name = ?", name).
		Where("type_id = ?", equipmentTypeID).
		Where("user_id = ?", userContext.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (r *pgEquipmentRepository) IsOwnedByUser(
	equipmentID uint,
	userContext authdom.UserContext,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("id = ?", equipmentID).
		Where("user_id = ?", userContext.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
