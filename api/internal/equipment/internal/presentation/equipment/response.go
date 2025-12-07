package http

import domain "octodome.com/api/internal/equipment/internal/domain/equipment"

type GetListResponse struct {
	Equipments []domain.EquipmentDTO `json:"equipments"`
	TotalCount int64                 `json:"total_count"`
}
