package mod

import (
	auth "octodome/internal/auth/internal/application"
	infra "octodome/internal/auth/internal/infrastructure"
	http "octodome/internal/auth/internal/presentation"
	userinfra "octodome/internal/user/infrastructure"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	ctrl := createAuthController(db)
	registerRoutes(r, ctrl)
}

func createAuthController(db *gorm.DB) *http.AuthController {
	repo := userinfra.NewPgUserRepository(db)
	tokenGenerator := infra.NewJwtTokenGenerator()
	passwordHasher := infra.NewBcryptPasswordHasher()
	authHandler := auth.NewAuthenticateHandler(repo, tokenGenerator, passwordHasher)

	return http.NewAuthController(authHandler)
}

func registerRoutes(r chi.Router, ctrl *http.AuthController) {
	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/", ctrl.Authenticate)
	})
}
