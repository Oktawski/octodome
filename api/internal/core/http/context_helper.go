package corehttp

import (
	"net/http"
	authdom "octodome.com/api/internal/auth/domain"
)

func GetUserContext(r *http.Request) (*authdom.UserContext, error) {
	user := r.Context().Value(authdom.UserContextKey)
	if user == nil {
		return nil, http.ErrNoCookie
	}

	userContext, ok := user.(*authdom.UserContext)
	if !ok {
		return nil, http.ErrNoCookie
	}

	return userContext, nil
}
