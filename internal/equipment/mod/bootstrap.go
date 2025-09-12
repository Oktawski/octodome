package mod

import (
	infra "octodome/internal/equipment/internal/infrastructure"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	registerEquipmentRoutes(r, createEquipmentController(db))
	registerEquipmentTypeRoutes(r, createEquipmentTypeController(db))
}
