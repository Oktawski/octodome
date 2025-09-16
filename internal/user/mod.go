package mod

import (
	authdom "octodome/internal/auth/domain"
	user "octodome/internal/user/internal/application"
	infra "octodome/internal/user/internal/infrastructure"
	http "octodome/internal/user/internal/presentation"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	registerRoutes(r, db, initializeController(db))
}

func initializeController(db *gorm.DB) *http.UserController {
	userRepo := infra.NewPgUserRepository(db)

	userCreateHandler := user.NewCreateHandler(userRepo)
	userGetByID := user.NewUserGetByIDHandler(userRepo)

	return http.NewUserController(userCreateHandler, userGetByID)
}

func registerRoutes(r chi.Router, db *gorm.DB, ctrl *http.UserController) {
	jwtMiddleware := middleware.JwtAuthMiddleware(db)

	r.Route("/user", func(user chi.Router) {

		user.Group(func(admin chi.Router) {
			admin.Use(jwtMiddleware)
			admin.Use(middleware.RequireRoles(authdom.RoleUser))

			admin.Post("/", ctrl.CreateUser)
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
