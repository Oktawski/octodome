package equipment

import (
	authdom "octodome/internal/auth/domain"
	eqdom "octodome/internal/equipment/domain"
)

type EquipmentTypeHandler interface {
	GetEquipmentTypes(user *authdom.UserContext) ([]EquipmentTypeGetResponse, error)
	GetEquipmentType(id uint, user *authdom.UserContext) (*EquipmentTypeGetResponse, error)
	CreateType(eqType *EquipmentTypeCreateCommand, userID uint) error
}

type equipmentTypeHandler struct {
	repo eqdom.EquipmentRepository
}

func NewEquipmentTypeHandler(repository eqdom.EquipmentRepository) EquipmentTypeHandler {
	return &equipmentTypeHandler{repo: repository}
}

func (h *equipmentTypeHandler) GetEquipmentTypes(
	user *authdom.UserContext) ([]EquipmentTypeGetResponse, error) {

	eqTypes, err := h.repo.GetEquipmentTypes(user)
	if err != nil {
		return nil, err
	}

	responses := make([]EquipmentTypeGetResponse, len(*eqTypes))
	for i, e := range *eqTypes {
		responses[i] = EquipmentTypeGetResponse{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return responses, nil
}

func (h *equipmentTypeHandler) GetEquipmentType(
	id uint,
	user *authdom.UserContext) (*EquipmentTypeGetResponse, error) {

	eqType, err := h.repo.GetEquipmentType(id, user)
	if err != nil {
		return nil, err
	}

	response := fromDomain(eqType)

	return response, nil
}

func (h *equipmentTypeHandler) CreateType(
	r *EquipmentTypeCreateCommand,
	userID uint) error {

	eqType := &eqdom.EquipmentType{
		Name:   r.Name,
		UserID: userID,
	}

	return h.repo.CreateType(eqType)
}
