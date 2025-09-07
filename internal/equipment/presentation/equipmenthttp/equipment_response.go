package eqhttp

import eqdom "octodome/internal/equipment/domain/equipment"

type GetListResponse struct {
	Equipments []eqdom.EquipmentDTO `json:"equipments"`
	TotalCount int64                `json:"total_count"`
}
