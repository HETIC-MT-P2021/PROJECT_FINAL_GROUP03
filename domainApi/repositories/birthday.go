package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
)

func FindAllBirthdays(birthdays *[]models.Birthday) error {
	return database.Db.Debug().Find(birthdays).Error
}

func FindBirthdayByServerID(b *models.Birthday) error {
	return database.Db.Debug().Where("server_id = ?", b.ServerID).First(&b).Error
}

func FindBirthdayByUserID(b *models.Birthday) error {
	return database.Db.Debug().Where("user_id = ?", b.UserID).First(&b).Error
}

func UpdateUserBirthday(b *models.Birthday) error {
	return database.Db.Debug().Model(&b).Where("user_id = ?", b.UserID).Update("message_sent", b.MessageSent).Error
}

func PersistBirthday(b *models.Birthday) error {
	return database.Db.Debug().Create(b).Error
}
