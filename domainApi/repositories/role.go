package repositories

import (
	"github.com/JackMaarek/go-bot-utils/database"
	"github.com/JackMaarek/go-bot-utils/models"
)

// FindAllRoles find all db roles
func FindAllRoles(roles *[]models.Role) error {
	return database.Db.Debug().Find(roles).Error
}

// FindRoleByName finds a role by its name
func FindRoleByName(r *models.Role) error {
	return database.Db.Debug().Model(&r).Where("name = ?", r.Name).First(&r).Error
}

// UpdateRoleByID updates a role
func UpdateRoleByID(r *models.Role) error {
	return database.Db.Debug().Model(&r).Where("id = ?", r.Id).Update("name", r.Name).Error
}

// DeleteRoleById deletes a role
func DeleteRoleById(r *models.Role) error {
	return database.Db.Debug().Model(&r).Delete(&r, r.Id).Error
}

// PersistRole creates a new role in the db
func PersistRole(r *models.Role) error {
	return database.Db.Debug().Create(r).Error
}
