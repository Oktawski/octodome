package http

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	cmd "octodome/internal/user/application/command"
	hdl "octodome/internal/user/application/handler"
	qry "octodome/internal/user/application/query"
)

type UserController struct {
	createHandler  *hdl.CreateHandler
	getByIDHandler *hdl.GetByIDHandler
}

func NewUserController(
	create *hdl.CreateHandler,
	getByID *hdl.GetByIDHandler) *UserController {
	return &UserController{
		createHandler:  create,
		getByIDHandler: getByID,
	}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil || id < 0 {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	query := qry.GetByID{ID: uint(id)}
	user, err := ctrl.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var command cmd.Create
	if err := corehttp.ParseJSON(r, &command); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: extend by ID
	corehttp.WriteJSON(w, http.StatusCreated, command)
}
