package mod

import (
	"octodome/internal/auth/domain"
	auth "octodome/internal/auth/internal/application"
	"octodome/internal/auth/internal/infrastructure/migration"
	"octodome/internal/auth/internal/infrastructure/repository"
	"octodome/internal/auth/internal/infrastructure/security"
	http "octodome/internal/auth/internal/presentation"
	userinfra "octodome/internal/user/infrastructure"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	migration.Migrate(db)
	ctrl := createAuthController(db)
	registerRoutes(r, db, ctrl)
}

func createAuthController(db *gorm.DB) *http.AuthController {
	userRepo := userinfra.NewPgUserRepository(db)
	authRepo := repository.NewPgRole(db)
	tokenGenerator := security.NewJwtTokenGenerator()
	passwordHasher := security.NewBcryptPasswordHasher()

	authenticateHandler := auth.NewAuthenticateHandler(userRepo, tokenGenerator, passwordHasher)
	assignRoleHandler := auth.NewAssignRoleHandler(authRepo)

	return http.NewAuthController(
		authenticateHandler,
		assignRoleHandler,
	)
}

func registerRoutes(r chi.Router, db *gorm.DB, ctrl *http.AuthController) {
	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/", ctrl.Authenticate)
	})

	r.Route("/auth/admin", func(authAdmin chi.Router) {
		authAdmin.Use(middleware.JwtAuthMiddleware(db))
		authAdmin.Use(middleware.RequireRoles(domain.RoleAdmin))

		authAdmin.Post("/assign-role", ctrl.AssignRole)
	})
}
