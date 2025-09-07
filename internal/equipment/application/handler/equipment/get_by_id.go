package hdl

import (
	qry "octodome/internal/equipment/application/query"
	domain "octodome/internal/equipment/domain/equipment"
)

type GetByIDHandler struct {
	repo domain.Repository
}

func NewGetByIDHandler(repo domain.Repository) *GetByIDHandler {
	return &GetByIDHandler{
		repo: repo,
	}
}

func (h *GetByIDHandler) Handle(q qry.EquipmentGetByID) (
	*domain.EquipmentDTO,
	error,
) {
	equipment, err := h.repo.GetByID(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &domain.EquipmentDTO{
		ID:   equipment.ID,
		Name: equipment.Name,
		Type: equipment.EquipmentType.ToDTO(),
	}

	return response, nil
}
