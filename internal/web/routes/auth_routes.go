package routes

import (
	authpres "octodome/internal/auth/presentation"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, controller *authpres.AuthController) {
	userGroup := r.Group("/auth")
	{
		userGroup.POST("", controller.Authenticate)
	}
}
