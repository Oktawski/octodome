package domain

import eqtypedom "octodome.com/api/internal/equipment/internal/domain/equipmenttype"

type EquipmentDTO struct {
	ID   uint
	Name string
	Type eqtypedom.EquipmentTypeDTO
}

type Equipment struct {
	ID              uint
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	EquipmentType   eqtypedom.EquipmentType
	UserID          uint
}
