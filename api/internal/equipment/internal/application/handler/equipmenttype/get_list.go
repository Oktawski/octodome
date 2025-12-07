package hdl

import (
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipmenttype"
)

type GetListHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewGetListHandler(deps dependencies.EquipmentTypeContainer) *GetListHandler {

	return &GetListHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
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
