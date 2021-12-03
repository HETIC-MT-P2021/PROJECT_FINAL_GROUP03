package repositories

import (
	"github.com/JackMaarek/go-bot-utils/database"
	"github.com/JackMaarek/go-bot-utils/models"
)

// PersistUser saves user in database
func PersistUser(u *models.User) error {
	return database.Db.Debug().Create(u).Error
}
