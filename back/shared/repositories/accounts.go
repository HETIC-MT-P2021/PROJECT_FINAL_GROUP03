package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
)

// FindAccountByDiscordID returns first user found with discord ID like account.DiscordID
func FindAccountByDiscordID(account *models.Account) error {
	return database.Db.Debug().Where("discord_id = ?", account.DiscordID).First(&account).Error
}

// PersistAccount persist account in database
func PersistAccount(account *models.Account) error {
	return database.Db.Debug().Create(account).Error
}

func UpdateAccount(account *models.Account) error {
	return database.Db.Debug().Save(&account).Error
}

func FindAccountServers(account *models.Account) ([]models.Server, error) {
	var servers []models.Server
	err := database.Db.Debug().Model(&account).Order("id ASC").Association("Servers").Find(&servers)

	return servers, err
}
