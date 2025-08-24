package eqinfra

import (
	"errors"
	authdom "octodome/internal/auth/domain"
	eqdom "octodome/internal/equipment/domain"

	"gorm.io/gorm"
)

type pgEquipmentTypeRepository struct {
	db *gorm.DB
}

func NewPgEquipmentTypeRepository(db *gorm.DB) *pgEquipmentTypeRepository {
	return &pgEquipmentTypeRepository{db: db}
}

func (r *pgEquipmentTypeRepository) GetEquipmentTypes(
	page int,
	pageSize int,
	user authdom.UserContext) ([]eqdom.EquipmentType, int64, error) {

	var eqTypes []equipmentType
	if dbError := r.db.
		Offset(page-1).
		Limit(pageSize).
		Where("user_id = ?", user.ID).
		Find(&eqTypes); dbError.Error != nil {
		return nil, 0, dbError.Error
	}

	domEqTypes := make([]eqdom.EquipmentType, len(eqTypes))
	for i, e := range eqTypes {
		domEqTypes[i] = eqdom.EquipmentType{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	var totalCount int64
	if err := r.db.
		Model(&equipmentType{}).
		Where("user_id = ?", user.ID).
		Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return domEqTypes, totalCount, nil
}

func (r *pgEquipmentTypeRepository) GetEquipmentType(
	id uint,
	user authdom.UserContext) (*eqdom.EquipmentType, error) {

	var eqType *equipmentType

	if dbError := r.db.
		Where("ID = ?", id).
		Where("user_id = ?", user.ID).
		First(&eqType).Error; dbError != nil {
		return nil, dbError
	}

	return eqType.toDomain(), nil
}

func (r *pgEquipmentTypeRepository) CreateType(eq *eqdom.EquipmentType) error {
	equipmentModel := fromDomain(eq)

	if err := r.db.Create(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentTypeRepository) DeleteEquipmentType(id uint, userContext authdom.UserContext) error {
	result := r.db.
		Where("id = ?", id).
		Where("user_id = ?", userContext.ID).
		Delete(&equipmentType{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("equipment type not found or not owned by the user")
	}

	return nil
}

func (r *pgEquipmentTypeRepository) ExistsByName(name string, userContext authdom.UserContext) bool {
	var count int64

	if err := r.db.Model(&equipmentType{}).
		Where("user_id = ?", userContext.ID).
		Where("name = ?", name).
		Count(&count).Error; err != nil {
		return false
	}

	return count == 0
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
