package http

import domain "octodome.com/api/internal/equipment/internal/domain/equipmenttype"

type GetListResponse struct {
	EquipmentTypes []domain.EquipmentTypeDTO `json:"equipmentTypes"`
	TotalCount     int64
}
