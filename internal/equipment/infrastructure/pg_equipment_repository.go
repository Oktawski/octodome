package eqinfra

import (
	authdom "octodome/internal/auth/domain"
	eqdom "octodome/internal/equipment/domain"

	"gorm.io/gorm"
)

type pgEquipmentRepository struct {
	db *gorm.DB
}

func NewPgEquipmentRepository(db *gorm.DB) *pgEquipmentRepository {
	return &pgEquipmentRepository{db: db}
}

func (r *pgEquipmentRepository) CreateType(eq *eqdom.EquipmentType) error {
	equipmentModel := fromDomain(eq)

	if err := r.db.Create(equipmentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentRepository) GetEquipmentTypes(
	page int,
	pageSize int,
	user *authdom.UserContext) ([]eqdom.EquipmentType, error) {

	var eqTypes []equipmentType
	if dbError := r.db.
		Offset(page-1).
		Limit(pageSize).
		Where("user_id = ?", user.ID).
		Find(&eqTypes); dbError != nil {
		return nil, dbError.Error
	}

	domEqTypes := make([]eqdom.EquipmentType, len(eqTypes))
	for i, e := range eqTypes {
		domEqTypes[i] = eqdom.EquipmentType{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return domEqTypes, nil
}

func (r *pgEquipmentRepository) GetEquipmentType(
	id uint,
	user *authdom.UserContext) (*eqdom.EquipmentType, error) {

	var eqType *equipmentType

	if dbError := r.db.
		Where("ID = ?", id).
		Where("user_id = ?", user.ID).
		First(&eqType).Error; dbError != nil {
		return nil, dbError
	}

	return eqType.toDomain(), nil
}
