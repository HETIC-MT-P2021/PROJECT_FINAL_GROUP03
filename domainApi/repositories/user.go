package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
)

func PersistUser(u *models.User) error {
	return database.Db.Debug().Create(u).Error
}
