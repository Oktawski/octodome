package hdl

import (
	qry "octodome/internal/equipment/application/query"
	domain "octodome/internal/equipment/domain/equipmenttype"
)

type GetByIDHandler struct {
	validator  domain.Validator
	repository domain.Repository
}

func NewGetByIDHandler(
	validator domain.Validator,
	repository domain.Repository) *GetByIDHandler {

	return &GetByIDHandler{
		validator:  validator,
		repository: repository,
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
