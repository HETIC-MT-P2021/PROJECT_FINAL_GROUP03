package repositories

import (
	"github.com/JackMaarek/go-bot-utils/database"
	"github.com/JackMaarek/go-bot-utils/models"
)

func FindAllRoles(roles *[]models.Role) error {
	return database.Db.Debug().Find(roles).Error
}

func FindRoleByName(r *models.Role) error {
	return database.Db.Debug().Model(&r).Where("name = ?", r.Name).First(&r).Error
}

func UpdateRoleByID(r *models.Role) error {
	return database.Db.Debug().Model(&r).Where("id = ?", r.Id).Update("name", r.Name).Error
}

func DeleteRoleById(r *models.Role) error {
	return database.Db.Debug().Model(&r).Delete(&r, r.Id).Error
}

func PersistRole(r *models.Role) error {
	return database.Db.Debug().Create(r).Error
}
