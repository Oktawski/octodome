package model

import (
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"

	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	EquipmentType   EquipmentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID          uint
}

func (Equipment) TableName() string {
	return "equipments"
}

func (e *Equipment) ToDomain() *domain.Equipment {
	return &domain.Equipment{
		ID:              e.ID,
		Name:            e.Name,
		Description:     e.Description,
		Category:        e.Category,
		EquipmentTypeID: e.EquipmentTypeID,
		UserID:          e.UserID,
	}
}

func EquipmentFromDomain(ed *domain.Equipment) *Equipment {
	return &Equipment{
		Model:           gorm.Model{ID: ed.ID},
		Name:            ed.Name,
		Description:     ed.Description,
		Category:        ed.Category,
		EquipmentTypeID: ed.EquipmentTypeID,
		UserID:          ed.UserID,
	}
}
