package eqtype

import (
	eqtypeqry "octodome/internal/equipment/application/query/equipmenttype"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
)

type GetListHandler struct {
	validator eqtypedom.Validator
	repo      eqtypedom.Repository
}

func NewGetListHandler(
	validator eqtypedom.Validator,
	repository eqtypedom.Repository) *GetListHandler {

	return &GetListHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *GetListHandler) Handle(q eqtypeqry.GetList) ([]eqtypedom.EquipmentTypeDTO, int64, error) {
	eqTypes, totalCount, err := h.repo.GetList(q.Pagination.Page, q.Pagination.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]eqtypedom.EquipmentTypeDTO, len(eqTypes))
	for i, e := range eqTypes {
		responses[i] = eqtypedom.EquipmentTypeDTO{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return responses, totalCount, nil
}
