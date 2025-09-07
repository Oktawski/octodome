package http

import domain "octodome/internal/equipment/domain/equipment"

type GetListResponse struct {
	Equipments []domain.EquipmentDTO `json:"equipments"`
	TotalCount int64                 `json:"total_count"`
}
