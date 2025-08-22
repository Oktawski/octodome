package eqinfra

import (
	eqdom "octodome/internal/equipment/domain"

	"gorm.io/gorm"
)

type equipmentType struct {
	gorm.Model
	Name   string
	UserID uint
}

func (equipmentType) TableName() string {
	return "equipment_types"
}

func (e *equipmentType) toDomain() *eqdom.EquipmentType {
	return &eqdom.EquipmentType{
		ID:     e.ID,
		Name:   e.Name,
		UserID: e.UserID,
	}
}

func fromDomain(e *eqdom.EquipmentType) *equipmentType {
	return &equipmentType{
		Model: gorm.Model{
			ID: e.ID,
		},
		Name:   e.Name,
		UserID: e.UserID,
	}
}
