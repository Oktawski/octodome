package mod

import (
	hdl "octodome/internal/equipment/internal/application/handler/equipmenttype"
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

	return http.NewEquipmentTypeController(
		hdl.NewCreateHandler(eqTypeValidator, eqTypeRepo),
		hdl.NewUpdateHandler(eqTypeValidator, eqTypeRepo),
		hdl.NewDeleteHandler(eqTypeValidator, eqTypeRepo),
		hdl.NewGetByIDHandler(eqTypeValidator, eqTypeRepo),
		hdl.NewGetListHandler(eqTypeValidator, eqTypeRepo))
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
