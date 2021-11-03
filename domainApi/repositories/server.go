package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
)

// FindServerByDiscordID returns first user found with discord ID like server.DiscordID
func FindServerByDiscordID(server *models.Server) error {
	return database.Db.Debug().Where("discord_id = ?", server.DiscordID).First(&server).Error
}

// PersistServer persist server in database
func PersistServer(server *models.Server) error {
	return database.Db.Debug().Create(server).Error
}

func UpdateServerMessage(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("welcome_message", server.WelcomeMessage).Error
}
func UpdateServerForbiddenWords(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("forbidden_words", server.ForbiddenWords).Error
}

func UpdateBirthdayMessage(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("birthday_message", server.BirthdayMessage).Error
}

func FindAllServers(servers *[]models.Server) error {
	return database.Db.Debug().Find(servers).Error
}
