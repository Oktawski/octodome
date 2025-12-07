package hdl

import (
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
)

type GetByIDHandler struct {
	repo domain.Repository
}

func NewGetByIDHandler(deps dependencies.EquipmentContainer) *GetByIDHandler {
	return &GetByIDHandler{
		repo: deps.Repository,
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
