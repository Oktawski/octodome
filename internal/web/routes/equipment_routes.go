package routes

import (
	eqpres "octodome/internal/equipment/presentation"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterEquipmentRoutes(
	r chi.Router,
	eqCtrl *eqpres.EquipmentController,
	eqTypeCtrl *eqpres.EquipmentTypeController) {

	r.Route("/equipment", func(eq chi.Router) {
		eq.Use(middleware.JwtAuthMiddleware)

		eq.Get("/", eqCtrl.GetEquipmentList)
		eq.Get("/{id:[0-9]+}", eqCtrl.GetEquipment)

		eq.Post("/", eqCtrl.CreateEquipment)
		eq.Put("/{id:[0-9]+}", eqCtrl.UpdateEquipment)
		eq.Delete("/{id:[0-9]+}", eqCtrl.DeleteEquipment)
	})

	r.Route("/equipment-type", func(eqType chi.Router) {
		eqType.Use(middleware.JwtAuthMiddleware)

		eqType.Get("/", eqTypeCtrl.GetEquipmentTypeList)
		eqType.Get("/{id:[0-9]+}", eqTypeCtrl.GetEquipmentType)

		eqType.Post("/", eqTypeCtrl.CreateEquipmentType)
		eqType.Put("/{id:[0-9]+}", eqTypeCtrl.UpdateEquipmentType)
		eqType.Delete("/{id:[0-9]+}", eqTypeCtrl.DeleteEquipmentType)
	})
}
