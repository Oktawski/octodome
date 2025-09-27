package mod

import (
	hdl "octodome/internal/equipment/internal/application/handler/equipmenttype"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
	repo "octodome/internal/equipment/internal/infrastructure/repository"
	http "octodome/internal/equipment/internal/presentation/equipmenttype"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func createEquipmentTypeController(db *gorm.DB) *http.EquipmentTypeController {
	eqTypeRepo := repo.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := domain.NewEquipmentTypeValidator(eqTypeRepo)

	deps := dependencies.NewEquipmentTypeContainer(eqTypeRepo, eqTypeValidator)

	return http.NewEquipmentTypeController(
		hdl.NewCreateHandler(deps),
		hdl.NewUpdateHandler(deps),
		hdl.NewDeleteHandler(deps),
		hdl.NewGetByIDHandler(deps),
		hdl.NewGetListHandler(deps))
}

func registerEquipmentTypeRoutes(r chi.Router, db *gorm.DB, ctrl *http.EquipmentTypeController) {
	r.Route("/equipment-type", func(eqType chi.Router) {
		eqType.Use(middleware.JwtAuthMiddleware(db))

		eqType.Get("/", ctrl.GetEquipmentTypeList)
		eqType.Get("/{id:[0-9]+}", ctrl.GetEquipmentType)

		eqType.Post("/", ctrl.CreateEquipmentType)
		eqType.Put("/{id:[0-9]+}", ctrl.UpdateEquipmentType)
		eqType.Delete("/{id:[0-9]+}", ctrl.DeleteEquipmentType)
	})
}
