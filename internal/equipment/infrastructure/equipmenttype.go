package eqinfra

import (
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"

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

func (e *equipmentType) toDomain() *eqtypedom.EquipmentType {
	return &eqtypedom.EquipmentType{
		ID:     e.ID,
		Name:   e.Name,
		UserID: e.UserID,
	}
}

func equipmentTypeFromDomain(e *eqtypedom.EquipmentType) *equipmentType {
	return &equipmentType{
		Model:  gorm.Model{ID: e.ID},
		Name:   e.Name,
		UserID: e.UserID,
	}
}
