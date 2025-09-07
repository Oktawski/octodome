package eqtype

import (
	eqtypeqry "octodome/internal/equipment/application/query/equipmenttype"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
)

type GetByIDHandler struct {
	validator  eqtypedom.Validator
	repository eqtypedom.Repository
}

func NewGetByIDHandler(
	validator eqtypedom.Validator,
	repository eqtypedom.Repository) *GetByIDHandler {

	return &GetByIDHandler{
		validator:  validator,
		repository: repository,
	}
}

func (h *GetByIDHandler) Handle(q eqtypeqry.GetByID) (*eqtypedom.EquipmentTypeDTO, error) {
	eqType, err := h.repository.GetByID(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &eqtypedom.EquipmentTypeDTO{
		ID:   eqType.ID,
		Name: eqType.Name,
	}

	return response, nil
}
