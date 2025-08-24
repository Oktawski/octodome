package userpres

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	user "octodome/internal/user/application"
)

type UserController struct {
	Handler user.UserHandler
}

func NewUserController(handler user.UserHandler) *UserController {
	return &UserController{Handler: handler}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil || id < 0 {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u, err := ctrl.Handler.GetUserByID(uint(id))
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, u)
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req user.UserCreateRequest

	if err := corehttp.ParseJSON(r, &req); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.Handler.CreateUser(&req); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusCreated, req)
}
