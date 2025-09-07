package http

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	cmd "octodome/internal/equipment/application/command"
	hdl "octodome/internal/equipment/application/handler/equipmenttype"
	qry "octodome/internal/equipment/application/query"
)

type EquipmentTypeController struct {
	createHandler  *hdl.CreateHandler
	updateHandler  *hdl.UpdateHandler
	deleteHandler  *hdl.DeleteHandler
	getByIDHandler *hdl.GetByIDHandler
	getListHandler *hdl.GetListHandler
}

func NewEquipmentTypeController(
	createHandler *hdl.CreateHandler,
	updateHandler *hdl.UpdateHandler,
	deleteHandler *hdl.DeleteHandler,
	getByIDHandler *hdl.GetByIDHandler,
	getListHandler *hdl.GetListHandler,
) *EquipmentTypeController {

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

	query := qry.EquipmentTypeGetList{Pagination: pagination, User: *user}
	eqTypes, totalCount, error := c.getListHandler.Handle(query)
	if error != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment types")
		return
	}

	response := &GetListResponse{EquipmentTypes: eqTypes, TotalCount: totalCount}
	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (c *EquipmentTypeController) GetEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	query := qry.EquipmentTypeGetByID{ID: uint(id), User: *user}
	eqType, err := c.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "equipment type not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, eqType)
}

func (c *EquipmentTypeController) CreateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var request CreateRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := cmd.EquipmentTypeCreate{Name: request.Name, UserContext: *user}
	if err := c.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: change response to include ID etc.
	corehttp.WriteJSON(w, http.StatusCreated, request)
}

func (c *EquipmentTypeController) UpdateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	var request UpdateRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	request.ID = uint(id)

	command := cmd.EquipmentTypeUpdate{ID: uint(id), Name: request.Name, UserContext: *user}
	if err := c.updateHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: return proper response
	corehttp.WriteJSON(w, http.StatusOK, request)
}

func (c *EquipmentTypeController) DeleteEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := cmd.EquipmentTypeDelete{ID: uint(id), UserContext: *user}
	if err := c.deleteHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
	}

	corehttp.WriteStatus(w, http.StatusOK)
}
