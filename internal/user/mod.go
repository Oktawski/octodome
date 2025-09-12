package mod

import (
	hdl "octodome/internal/user/internal/application/handler"
	infra "octodome/internal/user/internal/infrastructure"
	http "octodome/internal/user/internal/presentation"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	ctrl := initializeController(db)

	registerRoutes(r, ctrl)
}

func initializeController(db *gorm.DB) *http.UserController {
	userRepo := infra.NewPgUserRepository(db)

	userCreateHandler := hdl.NewCreateHandler(userRepo)
	userGetByID := hdl.NewUserGetByIDHandler(userRepo)

	return http.NewUserController(userCreateHandler, userGetByID)
}

func registerRoutes(r chi.Router, ctrl *http.UserController) {
	r.Route("/user", func(user chi.Router) {
		user.Post("/", ctrl.CreateUser)

		user.Group(func(protected chi.Router) {
			protected.Use(middleware.JwtAuthMiddleware)
			protected.Get("/{id:[0-9]+}", ctrl.GetUser)
		})
	})
}
