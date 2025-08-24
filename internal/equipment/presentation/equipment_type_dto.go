package eqpres

import eqdom "octodome/internal/equipment/domain"

type EquipmentTypeCreateDto struct {
	Name string `json:"name"`
}

type GetEquipmentTypesResponse struct {
	EqTypes    []eqdom.EquipmentTypeDTO `json:"equipmentTypes"`
	TotalCount int64
}
