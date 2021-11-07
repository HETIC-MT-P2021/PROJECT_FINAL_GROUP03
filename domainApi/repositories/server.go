package repositories

import (
	"github.com/JackMaarek/go-bot-utils/database"
	"github.com/JackMaarek/go-bot-utils/models"
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
