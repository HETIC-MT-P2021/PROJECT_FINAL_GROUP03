package services

import (
	"github.com/JackMaarek/go-bot-utils/enum"
	"github.com/JackMaarek/go-bot-utils/helpers"
	"github.com/JackMaarek/go-bot-utils/models"
	log "github.com/sirupsen/logrus"
)

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
