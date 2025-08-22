package eqinfra

import "gorm.io/gorm"

type equipment struct {
	gorm.Model
	TypeID uint
	Type   equipmentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
