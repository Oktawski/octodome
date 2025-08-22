package authpres

import (
	"net/http"
	auth "octodome/internal/auth/application"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Handler auth.AuthHandler
}

func NewAuthController(handler auth.AuthHandler) *AuthController {
	return &AuthController{Handler: handler}
}

func (ctrl *AuthController) Authenticate(c *gin.Context) {
	var authReq auth.AuthenticateRequest

	if err := c.ShouldBindJSON(&authReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.Handler.Authenticate(&authReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
