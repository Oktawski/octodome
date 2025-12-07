package mod

import (
	infra "octodome.com/api/internal/equipment/internal/infrastructure"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	registerEquipmentRoutes(r, db, createEquipmentController(db))
	registerEquipmentTypeRoutes(r, db, createEquipmentTypeController(db))
}
