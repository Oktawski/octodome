package http

import (
	"net/http"

	user "octodome.com/api/internal/user/internal/application"
	corehttp "octodome.com/shared/http"
)

type UserController struct {
	createHandler        *user.RegisterHandler
	getByIDHandler       *user.GetByIDHandler
	resetPasswordHandler *user.ResetPasswordHandler
}

func NewUserController(
	create *user.RegisterHandler,
	getByID *user.GetByIDHandler,
	resetPassword *user.ResetPasswordHandler,
) *UserController {
	return &UserController{
		createHandler:        create,
		getByIDHandler:       getByID,
		resetPasswordHandler: resetPassword,
	}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil || id < 0 {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	query := user.GetByID{Context: r.Context(), ID: uint(id)}
	user, err := ctrl.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, user)
}

func (ctrl *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var command user.Register
	command.Context = r.Context()
	if err := corehttp.ParseJSON(r, &command); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: extend by ID
	corehttp.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}

func (ctrl *UserController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var command user.ResetPassword
	command.Context = r.Context()
	if err := corehttp.ParseJSON(r, &command); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := ctrl.resetPasswordHandler.Handle(command)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, map[string]string{"message": "password reset"})
}
