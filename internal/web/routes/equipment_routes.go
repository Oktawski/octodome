package routes

import (
	eqpres "octodome/internal/equipment/presentation"
	"octodome/internal/web/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterEquipmentRoutes(r chi.Router, ctrl *eqpres.EquipmentController) {
	r.Route("/equipment-type", func(eqType chi.Router) {
		eqType.Use(middleware.JwtAuthMiddleware)

		eqType.Get("/", ctrl.GetEquipmentTypes)
		eqType.Get("/{id:[0-9]+}", ctrl.GetEquipmentType)

		eqType.Post("/", ctrl.CreateEquipmentType)

		eqType.Delete("/{id:[0-9]+}", ctrl.DeleteEquipmentType)
	})
}
