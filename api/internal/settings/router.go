package settings

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	settings "octodome.com/api/internal/settings/internal/application"
	infra "octodome.com/api/internal/settings/internal/infrastructure"
	"octodome.com/api/internal/settings/internal/presentation"
	"octodome.com/api/internal/web/middleware"

	authdom "octodome.com/api/internal/auth/domain"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	upsertHandler := settings.NewUpsertHandler(infra.NewPgSettingRepository(db))

	httpHandler := presentation.NewHttpHandler(upsertHandler)

	r.Route("/settings", func(settings chi.Router) {
		settings.Use(middleware.JwtAuthMiddleware(db))
		settings.Use(middleware.RequireAtLeastRole(authdom.RoleAdmin))

		settings.Post("/upsert", httpHandler.UpsertSetting)
	})
}
