package core

import (
	"errors"
	authdom "octodome/internal/auth/domain"

	"github.com/gin-gonic/gin"
)

func GetUserContext(c *gin.Context) (*authdom.UserContext, error) {
	user, exists := c.Get("userContext")

	if !exists {
		return nil, errors.New("user not found in context")
	}

	userContext, ok := user.(*authdom.UserContext)
	if !ok {
		return nil, errors.New("user in context has wrong type")
	}

	return userContext, nil
}
