package hdl

import (
	"context"

	corecontext "octodome.com/api/internal/core/context"
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
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
	ctx := context.WithValue(q.Ctx, corecontext.UserIDKey, q.UserContext.ID)
	equipments, totalCount, err := h.repo.GetList(ctx, q.Pagination.Page, q.Pagination.PageSize)
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
