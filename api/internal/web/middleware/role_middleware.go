package middleware

import (
	"fmt"
	"net/http"
	authdom "octodome.com/api/internal/auth/domain"
	"strings"
)

func RequireRoles(allowed ...authdom.RoleName) func(http.Handler) http.Handler {
	if len(allowed) == 0 {
		return func(next http.Handler) http.Handler { return next }
	}

	allowedSet := make(map[string]struct{}, len(allowed))
	for _, role := range allowed {
		allowedSet[strings.ToLower(string(role))] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxVal := r.Context().Value(authdom.UserContextKey)
			userCtx, ok := ctxVal.(*authdom.UserContext)
			if !ok || userCtx == nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, `{"error": "authentication required"}`)
				return
			}

			for _, role := range userCtx.Roles {
				if _, exists := allowedSet[strings.ToLower(string(role.Name))]; exists {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `{"error": "insufficient permissions"}`)
		})
	}
}

func RequireAtLeastRole(required authdom.RoleName) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxVal := r.Context().Value(authdom.UserContextKey)
			userCtx, ok := ctxVal.(*authdom.UserContext)
			if !ok || userCtx == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !userCtx.HasAtLeastRole(required) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
