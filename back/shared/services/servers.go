package services

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	log "github.com/sirupsen/logrus"
)

func ChangeWelcomeMessage(serverID, newMessage string) error {
	server := models.Server{DiscordID: serverID}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		log.Error(err)
		return err
	}

	// Update server
	server.WelcomeMessage = newMessage
	err := repositories.UpdateServerMessage(&server)
	if err != nil {
		log.Error(err)
	}

	return err
}
