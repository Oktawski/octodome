package eqpres

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	eqcommand "octodome/internal/equipment/application/command"
	eqhandler "octodome/internal/equipment/application/handler"
	eqquery "octodome/internal/equipment/application/query"
)

// package equipment

// import eqdom "octodome/internal/equipment/domain"

// type EquipmentTypeGetResponse struct {
// 	ID   uint
// 	Name string
// }

// func FromDomain(d *eqdom.EquipmentType) *EquipmentTypeGetResponse {
// 	return &EquipmentTypeGetResponse{
// 		ID:   d.ID,
// 		Name: d.Name,
// 	}
// }

type EquipmentController struct {
	eqHandler     eqhandler.EquipmentHandler
	eqTypeHandler eqhandler.EquipmentTypeHandler
}

func NewEquipmentController(
	eqHandler eqhandler.EquipmentHandler,
	eqTypeHandler eqhandler.EquipmentTypeHandler) *EquipmentController {

	return &EquipmentController{
		eqHandler:     eqHandler,
		eqTypeHandler: eqTypeHandler,
	}
}

func (ctrl *EquipmentController) GetEquipmentTypes(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	page := corehttp.GetQueryParamOrDefault(r, "page", 1)
	pageSize := corehttp.GetQueryParamOrDefault(r, "pageSize", 100)

	query := eqquery.GetListQuery{Page: page, PageSize: pageSize, User: *user}
	eqTypes, error := ctrl.eqTypeHandler.GetEquipmentTypes(query)
	if error != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment types")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, eqTypes)
}

func (ctrl *EquipmentController) GetEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	query := eqquery.GetQuery{ID: uint(id), User: *user}
	eqType, err := ctrl.eqTypeHandler.GetEquipmentType(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "equipment type not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, eqType)
}

func (ctrl *EquipmentController) CreateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var dto EquipmentTypeCreateDto
	if err := corehttp.ParseJSON(r, &dto); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := eqcommand.CreateCommand{Name: dto.Name, User: *user}
	if err := ctrl.eqTypeHandler.CreateType(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: change response to include ID etc.
	corehttp.WriteJSON(w, http.StatusCreated, dto)
}

func (ctrl *EquipmentController) DeleteEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := eqcommand.DeleteCommand{ID: uint(id), User: *user}
	if err := ctrl.eqTypeHandler.DeleteEquipmentType(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
	}

	corehttp.WriteStatus(w, http.StatusOK)
}
