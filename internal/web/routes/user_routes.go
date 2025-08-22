package routes

import (
	userpres "octodome/internal/user/presentation"
	"octodome/internal/web/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, controller *userpres.UserController) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("", controller.CreateUser)
	}

	userProtectedGroup := r.Group("/users")
	userProtectedGroup.Use(middleware.JwtAuthMiddleware())
	{
		userProtectedGroup.GET("/:id", controller.GetUser)
	}
}
