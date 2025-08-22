package equipment

import eqdom "octodome/internal/equipment/domain"

type EquipmentTypeGetResponse struct {
	ID   uint
	Name string
}

func fromDomain(d *eqdom.EquipmentType) *EquipmentTypeGetResponse {
	return &EquipmentTypeGetResponse{
		ID:   d.ID,
		Name: d.Name,
	}
}
