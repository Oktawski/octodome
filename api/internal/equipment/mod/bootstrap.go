package mod

import (
	"octodome.com/api/internal/equipment/internal/dependencies"
	equipmentdomain "octodome.com/api/internal/equipment/internal/domain/equipment"
	equipmenttypedomain "octodome.com/api/internal/equipment/internal/domain/equipmenttype"
	infra "octodome.com/api/internal/equipment/internal/infrastructure"
	repo "octodome.com/api/internal/equipment/internal/infrastructure/repository"
	equipmenthttp "octodome.com/api/internal/equipment/internal/presentation/equipment"
	equipmenttypehttp "octodome.com/api/internal/equipment/internal/presentation/equipmenttype"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Initialize(r chi.Router, db *gorm.DB) {
	infra.Migrate(db)

	registerEquipmentRoutes(r, db)
	registerEquipmentTypeRoutes(r, db)
}

func registerEquipmentRoutes(r chi.Router, db *gorm.DB) {
	eqRepo := repo.NewPgEquipmentRepository(db)
	equipmentValidator := equipmentdomain.NewValidator(eqRepo)

	deps := dependencies.NewEquipmentContainer(
		eqRepo,
		equipmentValidator,
	)

	equipmenthttp.RegisterEquipmentRoutes(r, db, deps)
}

func registerEquipmentTypeRoutes(r chi.Router, db *gorm.DB) {
	eqTypeRepo := repo.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := equipmenttypedomain.NewEquipmentTypeValidator(eqTypeRepo)

	deps := dependencies.NewEquipmentTypeContainer(eqTypeRepo, eqTypeValidator)

	equipmenttypehttp.RegisterEquipmentTypeRoutes(r, db, deps)
}
