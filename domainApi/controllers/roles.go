package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := repositories.FindAllRoles(&roles); err != nil {
		c.JSON(http.StatusNotFound, "Error while fetching roles")
		return
	}

	c.JSON(http.StatusOK, roles)
}

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

func UpdateRoleById(c *gin.Context) {
	var role models.Role
	role.Id = utils.ConvertStringToInt(c.Param("id"))
	if err := repositories.UpdateRoleByID(&role); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not update role with id %d", role.Id))
		return
	}

	c.JSON(http.StatusOK, "Role updated successfully")
}

func DeleteRoleById(c *gin.Context) {
	var role models.Role
	role.Id = utils.ConvertStringToInt(c.Param("id"))
	if err := repositories.DeleteRoleById(&role); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not delete role with id %d", role.Id))
		return
	}

	c.JSON(http.StatusOK, "Role deleted successfully")
}
