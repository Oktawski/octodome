package routes

import (
	userpres "octodome/internal/user/presentation"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, controller *userpres.UserController) {
	r.Route("/user", func(user chi.Router) {
		user.Post("/", controller.CreateUser)

		user.Group(func(protected chi.Router) {
			protected.Use(middleware.JwtAuthMiddleware)
			protected.Get("/{id:[0-9]+}", controller.GetUser)
		})
	})
}
