package authhttp

import (
	"net/http"
	auth "octodome/internal/auth/application"
	corehttp "octodome/internal/core/http"
)

type AuthController struct {
	AuthenticateHandler auth.AuthenticateHandler
}

func NewAuthController(handler auth.AuthenticateHandler) *AuthController {
	return &AuthController{AuthenticateHandler: handler}
}

func (ctrl *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var authCommand auth.AuthenticateCommand
	if err := corehttp.ParseJSON(r, &authCommand); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := ctrl.AuthenticateHandler.Handle(&authCommand)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, token)
}
