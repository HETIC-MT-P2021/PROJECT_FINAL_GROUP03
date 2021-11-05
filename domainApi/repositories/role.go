package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
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
