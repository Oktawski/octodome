package equipmenttype

import (
	eqquery "octodome/internal/equipment/application/query"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
)

type GetByIDHandler struct {
	validator  eqtypedom.EquipmentTypeValidator
	repository eqtypedom.EquipmentTypeRepository
}

func NewGetByIDHandler(
	validator eqtypedom.EquipmentTypeValidator,
	repository eqtypedom.EquipmentTypeRepository) *GetByIDHandler {

	return &GetByIDHandler{
		validator:  validator,
		repository: repository,
	}
}

func (h *GetByIDHandler) Handle(q eqquery.EquipmentTypeGetByID) (*eqtypedom.EquipmentTypeDTO, error) {
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
