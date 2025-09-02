package equipment

import (
	eqquery "octodome/internal/equipment/application/query"
	equipmentdom "octodome/internal/equipment/domain/equipment"
)

type GetListHandler struct {
	repo equipmentdom.EquipmentRepository
}

func NewGetListHandler(repo equipmentdom.EquipmentRepository) *GetListHandler {
	return &GetListHandler{
		repo: repo,
	}
}

func (h *GetListHandler) Handle(q eqquery.EquipmentGetList) ([]equipmentdom.EquipmentDTO, int64, error) {
	equipments, totalCount, err := h.repo.GetList(q.Pagination.Page, q.Pagination.PageSize, q.User)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]equipmentdom.EquipmentDTO, len(equipments))
	for i, e := range equipments {
		responses[i] = equipmentdom.EquipmentDTO{
			ID:   e.ID,
			Name: e.Name,
			Type: e.EquipmentType.ToDTO(),
		}
	}

	return responses, totalCount, nil
}
