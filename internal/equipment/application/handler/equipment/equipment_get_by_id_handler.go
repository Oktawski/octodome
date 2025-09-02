package equipment

import (
	eqquery "octodome/internal/equipment/application/query"
	equipmentdom "octodome/internal/equipment/domain/equipment"
)

type GetByIDHandler struct {
	repo equipmentdom.EquipmentRepository
}

func NewGetByIDHandler(repo equipmentdom.EquipmentRepository) *GetByIDHandler {
	return &GetByIDHandler{
		repo: repo,
	}
}

func (h *GetByIDHandler) Handle(q eqquery.EquipmentGetByID) (*equipmentdom.EquipmentDTO, error) {
	equipment, err := h.repo.GetByID(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &equipmentdom.EquipmentDTO{
		ID:   equipment.ID,
		Name: equipment.Name,
		Type: equipment.EquipmentType.ToDTO(),
	}

	return response, nil
}
