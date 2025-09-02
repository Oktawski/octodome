package eqpres

import eqtypedom "octodome/internal/equipment/domain/equipment_type"

type GetEquipmentTypesResponse struct {
	EqTypes    []eqtypedom.EquipmentTypeDTO `json:"equipmentTypes"`
	TotalCount int64
}

type EquipmentTypeCreateDto struct {
	Name string `json:"name"`
}

type EquipmentTypeUpdateDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
