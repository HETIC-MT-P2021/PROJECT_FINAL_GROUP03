package services

import (
	"github.com/JackMaarek/go-bot-utils/enum"
	"github.com/JackMaarek/go-bot-utils/helpers"
	"github.com/JackMaarek/go-bot-utils/models"
	log "github.com/sirupsen/logrus"
)

// GetRoles performs an http request to the discord api to get roles list
func GetRoles(roles *[]models.Role) error {
	response, err := helpers.PerformRequest(enum.RoleResourceRoute, "GET", "")
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while fetching roles")
	}
	err = helpers.DeserializeResponseFromObject(response.Body, roles)
	if err != nil {
		log.Error("could not deserialize roles object")
	}

	return nil
}

// UpdateRoleById performs an http request to the discord api to update a server role
func UpdateRoleById(id string, r *models.Role) error {
	url := enum.RoleResourceRoute + "/" + id
	response, err := helpers.PerformRequest(url, "PUT", r)
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while updating role")
	}

	return nil
}

// CreateRole performs an http request to the discord api to create a new role
func CreateRole(r *models.Role) error {
	response, err := helpers.PerformRequest(enum.RoleResourceRoute, "POST", r)
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while creating role")
	}

	return nil
}

// DeleteRole performs an http request to the discord api to delete a role
func DeleteRole(id string) error {
	url := enum.RoleResourceRoute + "/" + id
	response, err := helpers.PerformRequest(url, "DELETE", "")
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while updating role")
	}

	return nil
}
