package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	corehttp "octodome.com/api/internal/core/http"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	hdl "octodome.com/api/internal/equipment/internal/application/handler/equipmenttype"
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	"octodome.com/api/internal/web/middleware"
	sharedhttp "octodome.com/shared/http"
)

type equipmentTypeController struct {
	createHandler  *hdl.CreateHandler
	updateHandler  *hdl.UpdateHandler
	deleteHandler  *hdl.DeleteHandler
	getByIDHandler *hdl.GetByIDHandler
	getListHandler *hdl.GetListHandler
}

func RegisterEquipmentTypeRoutes(
	r chi.Router,
	db *gorm.DB,
	deps dependencies.EquipmentTypeContainer,
) {
	create := hdl.NewCreateHandler(deps)
	update := hdl.NewUpdateHandler(deps)
	delete := hdl.NewDeleteHandler(deps)
	getByID := hdl.NewGetByIDHandler(deps)
	getList := hdl.NewGetListHandler(deps)

	ctrl := &equipmentTypeController{
		createHandler:  create,
		updateHandler:  update,
		deleteHandler:  delete,
		getByIDHandler: getByID,
		getListHandler: getList,
	}

	r.Route("/equipment-type", func(eqType chi.Router) {
		eqType.Use(middleware.JwtAuthMiddleware(db))

		eqType.Get("/", ctrl.GetEquipmentTypeList)
		eqType.Get("/{id:[0-9]+}", ctrl.GetEquipmentType)

		eqType.Post("/", ctrl.CreateEquipmentType)
		eqType.Put("/{id:[0-9]+}", ctrl.UpdateEquipmentType)
		eqType.Delete("/{id:[0-9]+}", ctrl.DeleteEquipmentType)
	})
}

func (c *equipmentTypeController) GetEquipmentTypeList(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	pagination := sharedhttp.GetPagination(r)

	query := qry.EquipmentTypeGetList{User: *user, Pagination: pagination}
	eqTypes, totalCount, error := c.getListHandler.Handle(query)
	if error != nil {
		sharedhttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment types")
		return
	}

	response := &GetListResponse{EquipmentTypes: eqTypes, TotalCount: totalCount}
	sharedhttp.WriteJSON(w, http.StatusOK, response)
}

func (c *equipmentTypeController) GetEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := sharedhttp.GetPathParam[int](r, "id")
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	query := qry.EquipmentTypeGetByID{User: *user, ID: uint(id)}
	eqType, err := c.getByIDHandler.Handle(query)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "equipment type not found")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, eqType)
}

func (c *equipmentTypeController) CreateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var request CreateRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := cmd.EquipmentTypeCreate{Name: request.Name, UserContext: *user}
	if err := c.createHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: change response to include ID etc.
	sharedhttp.WriteJSON(w, http.StatusCreated, request)
}

func (c *equipmentTypeController) UpdateEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	id, err := sharedhttp.GetPathParam[int](r, "id")
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment type ID")
		return
	}

	var request UpdateRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	request.ID = uint(id)

	command := cmd.EquipmentTypeUpdate{
		UserContext: *user,
		Ctx:         r.Context(),
		ID:          uint(id),
		Name:        request.Name,
	}
	if err := c.updateHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	// TODO: return proper response
	sharedhttp.WriteJSON(w, http.StatusOK, request)
}

func (c *equipmentTypeController) DeleteEquipmentType(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	id, err := sharedhttp.GetPathParam[int](r, "id")
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	command := cmd.EquipmentTypeDelete{ID: uint(id), UserContext: *user}
	if err := c.deleteHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, err.Error())
	}

	sharedhttp.WriteStatus(w, http.StatusOK)
}
