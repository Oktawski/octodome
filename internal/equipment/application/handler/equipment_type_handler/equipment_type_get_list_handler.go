package equipmenttype

import (
	eqquery "octodome/internal/equipment/application/query"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
)

type GetListHandler struct {
	validator eqtypedom.EquipmentTypeValidator
	repo      eqtypedom.EquipmentTypeRepository
}

func NewGetListHandler(
	validator eqtypedom.EquipmentTypeValidator,
	repository eqtypedom.EquipmentTypeRepository) *GetListHandler {

	return &GetListHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *GetListHandler) Handle(q eqquery.EquipmentTypeGetList) ([]eqtypedom.EquipmentTypeDTO, int64, error) {
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
