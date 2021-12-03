package repositories

import (
	"github.com/JackMaarek/go-bot-utils/database"
	"github.com/JackMaarek/go-bot-utils/models"
)

// FindAllBirthdays returns all saved birthdays
func FindAllBirthdays(birthdays *[]models.Birthday) error {
	return database.Db.Debug().Find(birthdays).Error
}

// FindBirthdayByServerID finds all birthdays for a given server id
func FindBirthdayByServerID(b *models.Birthday) error {
	return database.Db.Debug().Where("server_id = ?", b.ServerID).First(&b).Error
}

// FindBirthdayByUserID finds a user birthday
func FindBirthdayByUserID(b *models.Birthday) error {
	return database.Db.Debug().Where("user_id = ?", b.UserID).First(&b).Error
}

// UpdateUserBirthday updates a user's birthday
func UpdateUserBirthday(b *models.Birthday) error {
	return database.Db.Debug().Model(&b).Where("user_id = ?", b.UserID).Update("message_sent", b.MessageSent).Error
}

// PersistBirthday registers a user's birthday
func PersistBirthday(b *models.Birthday) error {
	return database.Db.Debug().Create(b).Error
}
