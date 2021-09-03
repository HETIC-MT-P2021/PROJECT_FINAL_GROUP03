package services

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func ChangeWelcomeMessage(serverID, newMessage string, s *discordgo.Session) error {
	server := models.Server{DiscordID: serverID}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		log.Error(err)
		return err
	}

	// Update server
	server.WelcomeMessage = newMessage
	err := repositories.UpdateServer(&server)
	if err != nil {
		log.Error(err)
	}
	
	return err
}