package eqtypehandler

import (
	eqquery "octodome/internal/equipment/application/query"
	eqdom "octodome/internal/equipment/domain"
)

type GetListHandler struct {
	validator eqdom.EquipmentTypeValidator
	repo      eqdom.EquipmentTypeRepository
}

func NewGetListHandler(
	validator eqdom.EquipmentTypeValidator,
	repository eqdom.EquipmentTypeRepository) *GetListHandler {

	return &GetListHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *GetListHandler) Handle(q eqquery.GetList) ([]eqdom.EquipmentTypeDTO, int64, error) {
	eqTypes, totalCount, err := h.repo.GetEquipmentTypes(q.Page, q.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]eqdom.EquipmentTypeDTO, len(eqTypes))
	for i, e := range eqTypes {
		responses[i] = eqdom.EquipmentTypeDTO{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return responses, totalCount, nil
}
