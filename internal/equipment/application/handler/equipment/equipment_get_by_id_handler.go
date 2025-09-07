package eq

import (
	eqqry "octodome/internal/equipment/application/query/equipment"
	eqdom "octodome/internal/equipment/domain/equipment"
)

type GetByIDHandler struct {
	repo eqdom.Repository
}

func NewGetByIDHandler(repo eqdom.Repository) *GetByIDHandler {
	return &GetByIDHandler{
		repo: repo,
	}
}

func (h *GetByIDHandler) Handle(q eqqry.GetByID) (
	*eqdom.EquipmentDTO,
	error,
) {
	equipment, err := h.repo.GetByID(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &eqdom.EquipmentDTO{
		ID:   equipment.ID,
		Name: equipment.Name,
		Type: equipment.EquipmentType.ToDTO(),
	}

	return response, nil
}
