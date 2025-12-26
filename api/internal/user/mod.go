package mod

import (
	authdom "octodome.com/api/internal/auth/domain"
	user "octodome.com/api/internal/user/internal/application"
	infra "octodome.com/api/internal/user/internal/infrastructure"
	http "octodome.com/api/internal/user/internal/presentation"
	"octodome.com/api/internal/web/middleware"
	"octodome.com/shared/events"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	registerRoutes(r, db, initializeController(db))
}

func initializeController(db *gorm.DB) *http.UserController {
	userRepo := infra.NewPgUserRepository(db)
	eventsClient := events.NewClient("http://localhost:8990/events")

	userCreateHandler := user.NewCreateHandler(userRepo, *eventsClient)
	userGetByID := user.NewUserGetByIDHandler(userRepo)

	return http.NewUserController(userCreateHandler, userGetByID)
}

func registerRoutes(r chi.Router, db *gorm.DB, ctrl *http.UserController) {
	jwtMiddleware := middleware.JwtAuthMiddleware(db)

	r.Route("/user", func(user chi.Router) {
		// TODO: add rate limiting
		user.Post("/register", ctrl.CreateUser)

		user.Group(func(admin chi.Router) {
			admin.Use(jwtMiddleware)
			admin.Use(middleware.RequireRoles(authdom.RoleAdmin))
		})

		user.Group(func(atLeastUser chi.Router) {
			atLeastUser.Use(jwtMiddleware)
			atLeastUser.Use(middleware.RequireAtLeastRole(authdom.RoleUser))

			atLeastUser.Get("/{id:[0-9]+}", ctrl.GetUser)
		})

		user.Group(func(protected chi.Router) {
			protected.Use(middleware.JwtAuthMiddleware(db))
			// protected.Get("/{id:[0-9]+}", ctrl.GetUser)
		})
	})
}
