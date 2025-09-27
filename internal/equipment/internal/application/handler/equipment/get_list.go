package hdl

import (
	qry "octodome/internal/equipment/internal/application/query"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipment"
)

type GetListHandler struct {
	repo domain.Repository
}

func NewGetListHandler(deps dependencies.EquipmentContainer) *GetListHandler {
	return &GetListHandler{
		repo: deps.Repository,
	}
}

func (h *GetListHandler) Handle(q qry.EquipmentGetList) ([]domain.EquipmentDTO, int64, error) {
	equipments, totalCount, err := h.repo.GetList(q.Pagination.Page, q.Pagination.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]domain.EquipmentDTO, len(equipments))
	for i, e := range equipments {
		responses[i] = domain.EquipmentDTO{
			ID:   e.ID,
			Name: e.Name,
			Type: e.EquipmentType.ToDTO(),
		}
	}

	return responses, totalCount, nil
}
