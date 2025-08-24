package eqpres

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	eqcommand "octodome/internal/equipment/application/command"
	eqhandler "octodome/internal/equipment/application/handler"
	"octodome/internal/equipment/application/handler/eqtypehandler"
	eqquery "octodome/internal/equipment/application/query"
)

type EquipmentController struct {
	eqHandler            *eqhandler.EquipmentHandler
	eqTypeCreateHandler  *eqtypehandler.CreateHandler
	eqTypeDeleteHandler  *eqtypehandler.DeleteHandler
	eqTypeGetByIDHandler *eqtypehandler.GetByIDHandler
	eqTypeGetListHandler *eqtypehandler.GetListHandler
}

func NewEquipmentController(
	eqHandler *eqhandler.EquipmentHandler,
	CreateHandler *eqtypehandler.CreateHandler,
	DeleteHandler *eqtypehandler.DeleteHandler,
	GetByIDHandler *eqtypehandler.GetByIDHandler,
	getListHandler *eqtypehandler.GetListHandler) *EquipmentController {

	return &EquipmentController{
		eqHandler:            eqHandler,
		eqTypeCreateHandler:  CreateHandler,
		eqTypeDeleteHandler:  DeleteHandler,
		eqTypeGetByIDHandler: GetByIDHandler,
		eqTypeGetListHandler: getListHandler,
	}
}

func (ctrl *EquipmentController) GetEquipmentTypes(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	page := corehttp.GetQueryParamOrDefault(r, "page", 1)
	pageSize := corehttp.GetQueryParamOrDefault(r, "pageSize", 100)

	query := eqquery.GetList{Page: page, PageSize: pageSize, User: *user}
	eqTypes, totalCount, error := ctrl.eqTypeGetListHandler.Handle(query)
	if error != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment types")
		return
	}

	response := &GetEquipmentTypesResponse{EqTypes: eqTypes, TotalCount: totalCount}
	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (ctrl *EquipmentController) GetEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	query := eqquery.GetByID{ID: uint(id), User: *user}
	eqType, err := ctrl.eqTypeGetByIDHandler.Handle(query)
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
	if err := ctrl.eqTypeCreateHandler.Handle(command); err != nil {
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
	if err := ctrl.eqTypeDeleteHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
	}

	corehttp.WriteStatus(w, http.StatusOK)
}
