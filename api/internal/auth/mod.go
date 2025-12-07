package mod

import (
	auth "octodome.com/api/internal/auth/internal/application"
	"octodome.com/api/internal/auth/internal/dependencies"

	"octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/auth/internal/domain/validator"
	"octodome.com/api/internal/auth/internal/infrastructure/migration"
	"octodome.com/api/internal/auth/internal/infrastructure/repository"
	"octodome.com/api/internal/auth/internal/infrastructure/security"
	http "octodome.com/api/internal/auth/internal/presentation"
	userinfra "octodome.com/api/internal/user/infrastructure"
	"octodome.com/api/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	migration.Migrate(db)
	ctrl := createAuthController(db)
	registerRoutes(r, db, ctrl)
}

func createAuthController(db *gorm.DB) *http.AuthController {
	userReader := userinfra.NewPgUserRepository(db)
	roleRepo := repository.NewPgRole(db)
	tokenGenerator := security.NewJwtTokenGenerator()
	passwordHasher := security.NewBcryptPasswordHasher()
	validator := validator.NewRoleValidator(roleRepo)

	deps := dependencies.NewContainer(
		userReader,
		roleRepo,
		tokenGenerator,
		passwordHasher,
		validator,
	)

	authenticateHandler := auth.NewAuthenticateHandler(deps)
	assignRoleHandler := auth.NewAssignRoleHandler(deps)
	unassignRoleHandler := auth.NewUnassignRoleHandler(deps)
	syncRolesHandler := auth.NewSyncRolesHandler(deps)

	return http.NewAuthController(
		authenticateHandler,
		assignRoleHandler,
		unassignRoleHandler,
		syncRolesHandler,
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
		authAdmin.Post("/unassign-role", ctrl.UnassignRole)
		authAdmin.Post("/sync-roles", ctrl.SyncRoles)
	})
}
