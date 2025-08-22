package eqpres

import (
	"log"
	"net/http"
	"octodome/internal/core"
	equipment "octodome/internal/equipment/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EquipmentController struct {
	eqHandler     equipment.EquipmentHandler
	eqTypeHandler equipment.EquipmentTypeHandler
}

func NewEquipmentController(
	eqHandler equipment.EquipmentHandler,
	eqTypeHandler equipment.EquipmentTypeHandler) *EquipmentController {

	return &EquipmentController{
		eqHandler:     eqHandler,
		eqTypeHandler: eqTypeHandler,
	}
}

func (ctrl *EquipmentController) GetEquipmentTypes(c *gin.Context) {
	user, _ := core.GetUserContext(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	log.Print(page)
	log.Print(pageSize)

	eqTypes, error := ctrl.eqTypeHandler.GetEquipmentTypes(user)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not fetch equipment types"})
		return
	}

	c.JSON(http.StatusOK, eqTypes)

}

func (ctrl *EquipmentController) GetEquipmentType(c *gin.Context) {
	user, _ := core.GetUserContext(c)

	var idStr string = c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil || idInt < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid equipment type ID"})
		return
	}

	id := uint(idInt)

	eqType, err := ctrl.eqTypeHandler.GetEquipmentType(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment type not found"})
		return
	}

	c.JSON(http.StatusOK, eqType)
}

func (ctrl *EquipmentController) CreateEquipmentType(c *gin.Context) {
	user, _ := core.GetUserContext(c)

	var req equipment.EquipmentTypeCreateCommand
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.eqTypeHandler.CreateType(&req, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}
