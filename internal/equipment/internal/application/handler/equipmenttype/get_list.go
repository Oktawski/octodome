package hdl

import (
	qry "octodome/internal/equipment/internal/application/query"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
)

type GetListHandler struct {
	validator domain.Validator
	repo      domain.Repository
}

func NewGetListHandler(
	validator domain.Validator,
	repository domain.Repository) *GetListHandler {

	return &GetListHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *GetListHandler) Handle(q qry.EquipmentTypeGetList) ([]domain.EquipmentTypeDTO, int64, error) {
	eqTypes, totalCount, err := h.repo.GetList(q.Pagination.Page, q.Pagination.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]domain.EquipmentTypeDTO, len(eqTypes))
	for i, e := range eqTypes {
		responses[i] = domain.EquipmentTypeDTO{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return responses, totalCount, nil
}
