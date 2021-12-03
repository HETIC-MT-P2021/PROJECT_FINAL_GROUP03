package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/JackMaarek/go-bot-utils/models"
	"github.com/JackMaarek/go-bot-utils/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRoles is used to get all guild roles
func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := repositories.FindAllRoles(&roles); err != nil {
		c.JSON(http.StatusNotFound, "Error while fetching roles")
		return
	}

	c.JSON(http.StatusOK, roles)
}

// CreateRole is the controller called to create a new guild role
func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain valid properties")
		return
	}

	if err := repositories.PersistRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not create role %s", role.Name))
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Role %s created successfully", role.Name))
}

// UpdateRoleById is the controller used to update a guild role
func UpdateRoleById(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain valid properties")
		return
	}

	role.Id = utils.ConvertStringToInt(c.Param("id"))
	if err := repositories.UpdateRoleByID(&role); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not update role with id %d", role.Id))
		return
	}

	c.JSON(http.StatusOK, "Role updated successfully")
}

// DeleteRoleById is the controller used to delete one of the guild's role
func DeleteRoleById(c *gin.Context) {
	var role models.Role
	role.Id = utils.ConvertStringToInt(c.Param("id"))

	if err := repositories.DeleteRoleById(&role); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not delete role with id %d", role.Id))
		return
	}

	c.JSON(http.StatusOK, "Role deleted successfully")
}
