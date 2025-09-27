package repo

import (
	"errors"
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core/collection"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
	"octodome/internal/equipment/internal/infrastructure/model"

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
	user authdom.UserContext) ([]domain.EquipmentType, int64, error) {

	var eqTypes []model.EquipmentType
	var totalCount int64

	query := r.db.Model(&model.EquipmentType{}).Where("user_id = ?", user.ID)

	query.Count(&totalCount)
	query.Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&eqTypes).Error; err != nil {
		return nil, 0, err
	}

	domEqTypes := collection.Map(
		eqTypes,
		func(e model.EquipmentType) domain.EquipmentType {
			return *e.ToDomain()
		},
	)

	return domEqTypes, totalCount, nil
}

func (r *pgEquipmentTypeRepository) GetByID(
	id uint,
	user authdom.UserContext) (*domain.EquipmentType, error) {

	var eqType *model.EquipmentType

	if dbError := r.db.
		Where("ID = ?", id).
		Where("user_id = ?", user.ID).
		First(&eqType).Error; dbError != nil {
		return nil, dbError
	}

	return eqType.ToDomain(), nil
}

func (r *pgEquipmentTypeRepository) Create(eq *domain.EquipmentType) error {
	equipmentModel := model.EquipmentTypeFromDomain(eq)

	if err := r.db.Create(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentTypeRepository) Update(eq *domain.EquipmentType) error {
	equipmentModel := model.EquipmentTypeFromDomain(eq)

	if err := r.db.Save(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentTypeRepository) Delete(id uint) error {
	result := r.db.
		Where("id = ?", id).
		Delete(&model.EquipmentType{})

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
	userContext authdom.UserContext,
) bool {
	var count int64

	if err := r.db.Model(&model.EquipmentType{}).
		Where("user_id = ?", userContext.ID).
		Where("name = ?", name).
		Count(&count).Error; err != nil {
		return false
	}

	return count != 0
}

func (r *pgEquipmentTypeRepository) IsUsed(id uint, user authdom.UserContext) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
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

	if err := r.db.Model(&model.EquipmentType{}).
		Where("id = ?", id).
		Where("user_id = ?", user.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
