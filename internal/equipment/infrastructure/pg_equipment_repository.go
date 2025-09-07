package eqinfra

import (
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core/collection"
	eqdom "octodome/internal/equipment/domain/equipment"

	"gorm.io/gorm"
)

type pgEquipmentRepository struct {
	db *gorm.DB
}

func NewPgEquipmentRepository(db *gorm.DB) *pgEquipmentRepository {
	return &pgEquipmentRepository{db: db}
}

func (r *pgEquipmentRepository) GetList(
	page,
	pageSize int,
	user authdom.UserContext,
) ([]eqdom.Equipment, int64, error) {

	var equipments []equipment
	var total int64

	query := r.db.Model(&equipment{}).Where("user_id = ?", user.ID)

	query.Count(&total)
	query.Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&equipments).Error; err != nil {
		return nil, 0, err
	}

	equipmentList := collection.Map(
		equipments,
		func(e equipment) eqdom.Equipment {
			return *e.toDomain()
		},
	)

	return equipmentList, total, nil
}

func (r *pgEquipmentRepository) GetByID(
	id uint,
	user authdom.UserContext,
) (*eqdom.Equipment, error) {

	var eq *equipment

	if dbError := r.db.
		Where("ID = ?", id).
		Where("user_id = ?", user.ID).
		First(&eq).Error; dbError != nil {
		return nil, dbError
	}

	return eq.toDomain(), nil
}

func (r *pgEquipmentRepository) Create(e *eqdom.Equipment) error {
	equipment := equipmentFromDomain(e)

	if err := r.db.Create(equipment).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentRepository) Update(e *eqdom.Equipment) error {
	if dbError := r.db.
		Model(&equipment{}).
		Where("ID = ?", e.ID).
		Where("user_id = ?", e.UserID).
		Updates(equipmentFromDomain(e)).Error; dbError != nil {
		return dbError
	}
	return nil
}

func (r *pgEquipmentRepository) Delete(id uint) error {
	if dbError := r.db.
		Model(&equipment{}).
		Where("ID = ?", id).
		Delete(&equipment{}).Error; dbError != nil {
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

	if err := r.db.Model(&equipment{}).
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

	if err := r.db.Model(&equipment{}).
		Where("id = ?", equipmentID).
		Where("user_id = ?", userContext.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
