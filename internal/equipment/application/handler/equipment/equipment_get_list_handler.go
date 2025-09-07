package eq

import (
	eqqry "octodome/internal/equipment/application/query/equipment"
	eqdom "octodome/internal/equipment/domain/equipment"
)

type GetListHandler struct {
	repo eqdom.Repository
}

func NewGetListHandler(repo eqdom.Repository) *GetListHandler {
	return &GetListHandler{
		repo: repo,
	}
}

func (h *GetListHandler) Handle(q eqqry.GetList) ([]eqdom.EquipmentDTO, int64, error) {
	equipments, totalCount, err := h.repo.GetList(q.Pagination.Page, q.Pagination.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]eqdom.EquipmentDTO, len(equipments))
	for i, e := range equipments {
		responses[i] = eqdom.EquipmentDTO{
			ID:   e.ID,
			Name: e.Name,
			Type: e.EquipmentType.ToDTO(),
		}
	}

	return responses, totalCount, nil
}
