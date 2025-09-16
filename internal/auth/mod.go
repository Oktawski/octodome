package mod

import (
	auth "octodome/internal/auth/internal/application"
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/infrastructure/migration"
	"octodome/internal/auth/internal/infrastructure/repository"
	"octodome/internal/auth/internal/infrastructure/security"
	http "octodome/internal/auth/internal/presentation"
	userinfra "octodome/internal/user/infrastructure"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	migration.Migrate(db)
	ctrl := createAuthController(db)
	registerRoutes(r, ctrl)
}

func createAuthController(db *gorm.DB) *http.AuthController {
	repo := userinfra.NewPgUserRepository(db)
	tokenGenerator := security.NewJwtTokenGenerator()
	passwordHasher := security.NewBcryptPasswordHasher()
	authHandler := auth.NewAuthenticateHandler(repo, tokenGenerator, passwordHasher)

	return http.NewAuthController(authHandler)
}

func registerRoutes(r chi.Router, ctrl *http.AuthController) {
	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/", ctrl.Authenticate)
	})
}

func ProvideRoleReader(db *gorm.DB) domain.RoleRepository {
	return repository.NewPgRole(db)
}
