package eqtypehttp

import eqtypedom "octodome/internal/equipment/domain/equipmenttype"

type GetEquipmentTypesResponse struct {
	EquipmentTypes []eqtypedom.EquipmentTypeDTO `json:"equipmentTypes"`
	TotalCount     int64
}
