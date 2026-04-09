package mod

import (
	auth "octodome.com/api/internal/auth/internal/application"
	"octodome.com/api/internal/auth/internal/dependencies"
	"octodome.com/shared/events"

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
	ctrl := createAuthenticateHandlers(db)
	registerRoutes(r, db, ctrl)
}

func createAuthenticateHandlers(db *gorm.DB) *http.AuthHandlers {
	userReader := userinfra.NewPgUserRepository(db)
	roleRepo := repository.NewPgRole(db)
	tokenGenerator := security.NewJwtTokenGenerator()
	passwordHasher := security.NewBcryptPasswordHasher()
	validator := validator.NewRoleValidator(roleRepo)
	magicCodeRepo := repository.NewPgMagicCode(db)
	eventsClient := events.NewClient("http://event-broker:8990")

	deps := dependencies.NewContainer(
		userReader,
		roleRepo,
		tokenGenerator,
		passwordHasher,
		validator,
		magicCodeRepo,
		eventsClient,
	)

	authenticateHandler := auth.NewAuthenticateCredentialsHandler(deps)
	sendMagicCodeHandler := auth.NewSendMagicCodeHandler(deps)
	authenticateMagicCodeHandler := auth.NewAuthenticateMagicCodeHandler(deps)
	assignRoleHandler := auth.NewAssignRoleHandler(deps)
	unassignRoleHandler := auth.NewUnassignRoleHandler(deps)
	syncRolesHandler := auth.NewSyncRolesHandler(deps)

	return http.NewAuthHandlers(
		authenticateHandler,
		sendMagicCodeHandler,
		authenticateMagicCodeHandler,
		assignRoleHandler,
		unassignRoleHandler,
		syncRolesHandler,
	)
}

func registerRoutes(r chi.Router, db *gorm.DB, handlers *http.AuthHandlers) {
	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/credentials", handlers.AuthenticateCredentials)
		auth.Post("/send-magic-code", handlers.SendMagicCode)
		auth.Post("/authenticate-magic-code", handlers.AuthenticateMagicCode)
	})

	r.Route("/auth/admin", func(authAdmin chi.Router) {
		authAdmin.Use(middleware.JwtAuthMiddleware(db))
		authAdmin.Use(middleware.RequireRoles(domain.RoleAdmin))

		authAdmin.Post("/assign-role", handlers.AssignRole)
		authAdmin.Post("/unassign-role", handlers.UnassignRole)
		authAdmin.Post("/sync-roles", handlers.SyncRoles)
	})
}
