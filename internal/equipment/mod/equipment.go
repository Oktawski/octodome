package mod

import (
	hdl "octodome/internal/equipment/internal/application/handler/equipment"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipment"
	repo "octodome/internal/equipment/internal/infrastructure/repository"
	http "octodome/internal/equipment/internal/presentation/equipment"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func createEquipmentController(db *gorm.DB) *http.EquipmentController {
	eqRepo := repo.NewPgEquipmentRepository(db)
	equipmentValidator := domain.NewValidator(eqRepo)

	deps := dependencies.NewEquipmentContainer(
		eqRepo,
		equipmentValidator,
	)

	create := hdl.NewCreateHandler(deps)
	update := hdl.NewUpdateHandler(deps)
	delete := hdl.NewDeleteHandler(deps)
	getByID := hdl.NewGetByIDHandler(deps)
	getList := hdl.NewGetListHandler(deps)

	return http.NewEquipmentController(
		create,
		update,
		delete,
		getByID,
		getList,
	)
}

func registerEquipmentRoutes(r chi.Router, db *gorm.DB, ctrl *http.EquipmentController) {
	r.Route("/equipment", func(eq chi.Router) {
		eq.Use(middleware.JwtAuthMiddleware(db))

		eq.Get("/", ctrl.GetEquipmentList)
		eq.Get("/{id:[0-9]+}", ctrl.GetEquipment)

		eq.Post("/", ctrl.CreateEquipment)
		eq.Put("/{id:[0-9]+}", ctrl.UpdateEquipment)
		eq.Delete("/{id:[0-9]+}", ctrl.DeleteEquipment)
	})
}
