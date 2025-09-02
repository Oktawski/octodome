package eqpres

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	eqcommand "octodome/internal/equipment/application/command"
	equipmenttype "octodome/internal/equipment/application/handler/equipment_type_handler"
	eqquery "octodome/internal/equipment/application/query"
)

type EquipmentTypeController struct {
	createHandler  *equipmenttype.CreateHandler
	updateHandler  *equipmenttype.UpdateHandler
	deleteHandler  *equipmenttype.DeleteHandler
	getByIDHandler *equipmenttype.GetByIDHandler
	getListHandler *equipmenttype.GetListHandler
}

func NewEquipmentTypeController(
	createHandler *equipmenttype.CreateHandler,
	updateHandler *equipmenttype.UpdateHandler,
	deleteHandler *equipmenttype.DeleteHandler,
	getByIDHandler *equipmenttype.GetByIDHandler,
	getListHandler *equipmenttype.GetListHandler) *EquipmentTypeController {

	return &EquipmentTypeController{
		createHandler:  createHandler,
		updateHandler:  updateHandler,
		deleteHandler:  deleteHandler,
		getByIDHandler: getByIDHandler,
		getListHandler: getListHandler,
	}
}

func (c *EquipmentTypeController) GetEquipmentTypeList(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	pagination := corehttp.GetPagination(r)

	query := eqquery.EquipmentTypeGetList{Pagination: pagination, User: *user}
	eqTypes, totalCount, error := c.getListHandler.Handle(query)
	if error != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment types")
		return
	}

	response := &GetEquipmentTypesResponse{EqTypes: eqTypes, TotalCount: totalCount}
	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (c *EquipmentTypeController) GetEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	query := eqquery.EquipmentTypeGetByID{ID: uint(id), User: *user}
	eqType, err := c.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "equipment type not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, eqType)
}

func (c *EquipmentTypeController) CreateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var dto EquipmentTypeCreateDto
	if err := corehttp.ParseJSON(r, &dto); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := eqcommand.EquipmentTypeCreateCommand{Name: dto.Name, UserContext: *user}
	if err := c.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: change response to include ID etc.
	corehttp.WriteJSON(w, http.StatusCreated, dto)
}

func (c *EquipmentTypeController) UpdateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	var dto EquipmentTypeUpdateDTO
	if err := corehttp.ParseJSON(r, &dto); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	dto.ID = uint(id)

	command := eqcommand.EquipmentTypeUpdateCommand{ID: uint(id), Name: dto.Name, UserContext: *user}
	if err := c.updateHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, dto)
}

func (c *EquipmentTypeController) DeleteEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := eqcommand.EquipmentTypeDeleteCommand{ID: uint(id), UserContext: *user}
	if err := c.deleteHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
	}

	corehttp.WriteStatus(w, http.StatusOK)
}
