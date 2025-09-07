package eqinfra

import (
	eqdom "octodome/internal/equipment/domain/equipment"

	"gorm.io/gorm"
)

type equipment struct {
	gorm.Model
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	EquipmentType   equipmentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID          uint
}

func (equipment) TableName() string {
	return "equipments"
}

func (e *equipment) toDomain() *eqdom.Equipment {
	return &eqdom.Equipment{
		ID:              e.ID,
		Name:            e.Name,
		Description:     e.Description,
		Category:        e.Category,
		EquipmentTypeID: e.EquipmentTypeID,
		UserID:          e.UserID,
	}
}

func equipmentFromDomain(ed *eqdom.Equipment) *equipment {
	return &equipment{
		Model:           gorm.Model{ID: ed.ID},
		Name:            ed.Name,
		Description:     ed.Description,
		Category:        ed.Category,
		EquipmentTypeID: ed.EquipmentTypeID,
		UserID:          ed.UserID,
	}
}
