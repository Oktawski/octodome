package routes

import (
	authpres "octodome/internal/auth/presentation"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, controller *authpres.AuthController) {
	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/", controller.Authenticate)
	})
}
