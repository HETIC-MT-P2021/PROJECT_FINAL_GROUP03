package services

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/enum"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	log "github.com/sirupsen/logrus"
)

func GetRoles(roles *[]models.Role) error {
	response, err := PerformApiRequest(enum.RolesResourceRoute, "GET", "")
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while fetching roles")
	}
	err = DeserializeResponseFromObject(response.Body, roles)
	if err != nil {
		log.Error("could not deserialize roles object")
	}

	return nil
}

func UpdateRoleById(id string, r *models.Role) error {
	url := enum.RolesResourceRoute + "/" + id
	response, err := PerformApiRequest(url, "PUT", r)
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while updating role")
	}

	return nil
}

func CreateRole(r *models.Role) error {
	response, err := PerformApiRequest(enum.RolesResourceRoute, "POST", r)
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while creating role")
	}

	return nil
}

func DeleteRole(id string) error {
	url := enum.RolesResourceRoute + "/" + id
	response, err := PerformApiRequest(url, "DELETE", "")
	if err != nil {
		log.Error(err)
	}
	if response.StatusCode != 200 {
		log.Error("Error occured while updating role")
	}

	return nil
}
