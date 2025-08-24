package eqhandler

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqquery "octodome/internal/equipment/application/query"
	eqdom "octodome/internal/equipment/domain"
)

type EquipmentTypeHandler interface {
	GetEquipmentTypes(q eqquery.GetListQuery) ([]eqdom.EquipmentTypeDto, error)
	GetEquipmentType(q eqquery.GetQuery) (*eqdom.EquipmentTypeDto, error)
	CreateType(c eqcommand.CreateCommand) error
	DeleteEquipmentType(c eqcommand.DeleteCommand) error
}

type equipmentTypeHandler struct {
	validator eqdom.EquipmentTypeValidator
	repo      eqdom.EquipmentTypeRepository
}

func NewEquipmentTypeHandler(
	validator eqdom.EquipmentTypeValidator,
	repository eqdom.EquipmentTypeRepository) EquipmentTypeHandler {

	return &equipmentTypeHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *equipmentTypeHandler) GetEquipmentTypes(q eqquery.GetListQuery) ([]eqdom.EquipmentTypeDto, error) {

	eqTypes, err := h.repo.GetEquipmentTypes(q.Page, q.PageSize, q.User)
	if err != nil {
		return nil, err
	}

	responses := make([]eqdom.EquipmentTypeDto, len(eqTypes))
	for i, e := range eqTypes {
		responses[i] = eqdom.EquipmentTypeDto{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	return responses, nil
}

func (h *equipmentTypeHandler) GetEquipmentType(q eqquery.GetQuery) (*eqdom.EquipmentTypeDto, error) {
	eqType, err := h.repo.GetEquipmentType(q.ID, q.User)
	if err != nil {
		return nil, err
	}

	response := &eqdom.EquipmentTypeDto{
		ID:   eqType.ID,
		Name: eqType.Name,
	}

	return response, nil
}

func (h *equipmentTypeHandler) CreateType(c eqcommand.CreateCommand) error {
	if !h.validator.CanBeCreated(c.Name, c.User) {
		return errors.New("equipment type with this name already exists")
	}

	eqType := &eqdom.EquipmentType{
		Name:   c.Name,
		UserID: c.User.ID,
	}

	return h.repo.CreateType(eqType)
}

func (h *equipmentTypeHandler) DeleteEquipmentType(c eqcommand.DeleteCommand) error {

	if !h.validator.CanBeDeleted(c.ID, c.User) {
		return errors.New("equipment type cannot be removed")
	}

	if err := h.repo.DeleteEquipmentType(c.ID, c.User); err != nil {
		return err
	}

	return nil
}
