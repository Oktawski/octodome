package eqtypehandler

import (
	eqquery "octodome/internal/equipment/application/query"
	eqdom "octodome/internal/equipment/domain"
)

type GetByIDHandler struct {
	validator  eqdom.EquipmentTypeValidator
	repository eqdom.EquipmentTypeRepository
}

func NewGetByIDHandler(
	validator eqdom.EquipmentTypeValidator,
	repository eqdom.EquipmentTypeRepository) *GetByIDHandler {

	return &GetByIDHandler{
		validator:  validator,
		repository: repository,
	}
}

func (h *GetByIDHandler) Handle(q eqquery.GetByID) (*eqdom.EquipmentTypeDTO, error) {
	eqType, err := h.repository.GetEquipmentType(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &eqdom.EquipmentTypeDTO{
		ID:   eqType.ID,
		Name: eqType.Name,
	}

	return response, nil
}
