package model

import (
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
	changetracker "octodome.com/shared/changetracker"

	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Version         uint
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

func (e *Equipment) Update(userID uint, ed *domain.Equipment) *changetracker.ChangeTracker {
	ct := changetracker.New(userID)

	changetracker.UpdateWhenNotEqual(ct,
		func() string { return e.Name },
		func(v string) { e.Name = v },
		ed.Name,
		"name")

	changetracker.UpdateWhenNotEqual(ct,
		func() string { return e.Description },
		func(v string) { e.Description = v },
		ed.Description,
		"description")

	changetracker.UpdateWhenNotEqual(ct,
		func() string { return e.Category },
		func(v string) { e.Category = v },
		ed.Category,
		"category")

	changetracker.UpdateWhen(ct,
		ed.EquipmentTypeID != 0,
		func() uint { return e.EquipmentTypeID },
		func(v uint) { e.EquipmentTypeID = v },
		ed.EquipmentTypeID,
		"equipment_type_id")

	if ct.HasChanges {
		e.Version++
	}

	return ct
}
