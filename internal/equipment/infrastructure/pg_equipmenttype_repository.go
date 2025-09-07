package eqinfra

import (
	"errors"
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core/collection"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"

	"gorm.io/gorm"
)

type pgEquipmentTypeRepository struct {
	db *gorm.DB
}

func NewPgEquipmentTypeRepository(db *gorm.DB) *pgEquipmentTypeRepository {
	return &pgEquipmentTypeRepository{db: db}
}

func (r *pgEquipmentTypeRepository) GetList(
	page int,
	pageSize int,
	user authdom.UserContext) ([]eqtypedom.EquipmentType, int64, error) {

	var eqTypes []equipmentType
	var totalCount int64

	query := r.db.Model(&equipmentType{}).Where("user_id = ?", user.ID)

	query.Count(&totalCount)
	query.Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&eqTypes).Error; err != nil {
		return nil, 0, err
	}

	domEqTypes := collection.Map(
		eqTypes,
		func(e equipmentType) eqtypedom.EquipmentType {
			return *e.toDomain()
		},
	)

	return domEqTypes, totalCount, nil
}

func (r *pgEquipmentTypeRepository) GetByID(
	id uint,
	user authdom.UserContext) (*eqtypedom.EquipmentType, error) {

	var eqType *equipmentType

	if dbError := r.db.
		Where("ID = ?", id).
		Where("user_id = ?", user.ID).
		First(&eqType).Error; dbError != nil {
		return nil, dbError
	}

	return eqType.toDomain(), nil
}

func (r *pgEquipmentTypeRepository) Create(eq *eqtypedom.EquipmentType) error {
	equipmentModel := equipmentTypeFromDomain(eq)

	if err := r.db.Create(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentTypeRepository) Update(eq *eqtypedom.EquipmentType) error {
	equipmentModel := equipmentTypeFromDomain(eq)

	if err := r.db.Save(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentTypeRepository) Delete(id uint) error {
	result := r.db.
		Where("id = ?", id).
		Delete(&equipmentType{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("equipment type not found or not owned by the user")
	}

	return nil
}

func (r *pgEquipmentTypeRepository) ExistsByName(
	name string,
	userContext authdom.UserContext) bool {

	var count int64

	if err := r.db.Model(&equipmentType{}).
		Where("user_id = ?", userContext.ID).
		Where("name = ?", name).
		Count(&count).Error; err != nil {
		return false
	}

	return count != 0
}

func (r *pgEquipmentTypeRepository) IsUsed(id uint, user authdom.UserContext) bool {
	var count int64

	if err := r.db.Model(&equipment{}).
		Joins("JOIN equipment_types et ON equipment.type_id = et.id").
		Where("type_id = ?", id).
		Where("et.user_id = ?", user.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (r *pgEquipmentTypeRepository) OwnedByUser(id uint, user authdom.UserContext) bool {
	var count int64

	if err := r.db.Model(&equipmentType{}).
		Where("id = ?", id).
		Where("user_id = ?", user.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
