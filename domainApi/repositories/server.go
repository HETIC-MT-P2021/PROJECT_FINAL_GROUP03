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

// UpdateServerMessage updates the welcome message for a server
func UpdateServerMessage(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("welcome_message", server.WelcomeMessage).Error
}

// UpdateServerForbiddenWords updates the forbidden words list of a server
func UpdateServerForbiddenWords(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("forbidden_words", server.ForbiddenWords).Error
}

// UpdateBirthdayMessage updtaes the messages sent on birthdays for a server
func UpdateBirthdayMessage(server *models.Server) error {
	return database.Db.Debug().Model(&server).Where("discord_id = ?", server.DiscordID).Update("birthday_message", server.BirthdayMessage).Error
}

// FindAllServers find all bot servers
func FindAllServers(servers *[]models.Server) error {
	return database.Db.Debug().Find(servers).Error
}
