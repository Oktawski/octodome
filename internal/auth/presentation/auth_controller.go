package authpres

import (
	"net/http"
	auth "octodome/internal/auth/application"
	corehttp "octodome/internal/core/http"
)

type AuthController struct {
	Handler auth.AuthHandler
}

func NewAuthController(handler auth.AuthHandler) *AuthController {
	return &AuthController{Handler: handler}
}

func (ctrl *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var authReq auth.AuthenticateRequest

	if err := corehttp.ParseJSON(r, &authReq); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := ctrl.Handler.Authenticate(&authReq)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, token)
}
