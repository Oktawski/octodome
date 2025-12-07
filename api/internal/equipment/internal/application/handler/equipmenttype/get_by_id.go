package hdl

import (
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipmenttype"
)

type GetByIDHandler struct {
	repository domain.Repository
	validator  domain.Validator
}

func NewGetByIDHandler(deps dependencies.EquipmentTypeContainer) *GetByIDHandler {

	return &GetByIDHandler{
		repository: deps.Repository,
		validator:  deps.Validator,
	}
}

func (h *GetByIDHandler) Handle(q qry.EquipmentTypeGetByID) (*domain.EquipmentTypeDTO, error) {
	eqType, err := h.repository.GetByID(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &domain.EquipmentTypeDTO{
		ID:   eqType.ID,
		Name: eqType.Name,
	}

	return response, nil
}
