package model

import (
	domain "octodome/internal/equipment/internal/domain/equipmenttype"

	"gorm.io/gorm"
)

type EquipmentType struct {
	gorm.Model
	Name   string
	UserID uint
}

func (EquipmentType) TableName() string {
	return "equipment_types"
}

func (e *EquipmentType) ToDomain() *domain.EquipmentType {
	return &domain.EquipmentType{
		ID:     e.ID,
		Name:   e.Name,
		UserID: e.UserID,
	}
}

func EquipmentTypeFromDomain(e *domain.EquipmentType) *EquipmentType {
	return &EquipmentType{
		Model:  gorm.Model{ID: e.ID},
		Name:   e.Name,
		UserID: e.UserID,
	}
}
