package eqpres

import equipmentdom "octodome/internal/equipment/domain/equipment"

type GetEquipmentListResponse struct {
	Equipments []equipmentdom.EquipmentDTO `json:"equipments"`
	TotalCount int64                       `json:"total_count"`
}

type EquipmentCreateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type EquipmentUpdateDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
