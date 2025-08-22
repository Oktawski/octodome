package userpres

import (
	"net/http"
	user "octodome/internal/user/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Handler user.UserHandler
}

func NewUserController(handler user.UserHandler) *UserController {
	return &UserController{Handler: handler}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	var idStr string = c.Param("id")

	idInt, err := strconv.Atoi(idStr)
	if err != nil || idInt < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	id := uint(idInt)

	user, err := ctrl.Handler.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user user.UserCreateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Handler.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
