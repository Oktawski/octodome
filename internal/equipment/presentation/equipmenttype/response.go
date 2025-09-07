package http

import domain "octodome/internal/equipment/domain/equipmenttype"

type GetListResponse struct {
	EquipmentTypes []domain.EquipmentTypeDTO `json:"equipmentTypes"`
	TotalCount     int64
}
