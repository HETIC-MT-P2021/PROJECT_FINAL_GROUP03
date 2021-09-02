package repositories

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
)

// FindServerByDiscordID returns first user found with discord ID like server.DiscordID
func FindServerByDiscordID(server *models.Server) error {
	return database.Db.Debug().First(&server).Where("discord_id = ?", server.DiscordID).Error
}

// PersistServer persist server in database
func PersistServer(server *models.Server) error {
	return database.Db.Debug().Create(server).Error
}