package routes

import (
	eqpres "octodome/internal/equipment/presentation"
	"octodome/internal/web/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEquipmentRoutes(r *gin.Engine, ctrl *eqpres.EquipmentController) {
	eqTypes := r.Group("/equipment-type")
	eqTypes.Use(middleware.JwtAuthMiddleware())
	{
		eqTypes.GET("", ctrl.GetEquipmentTypes)
		eqTypes.GET("/:id", ctrl.GetEquipmentType)
		eqTypes.POST("", ctrl.CreateEquipmentType)
	}
}
