package domain

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	log "github.com/sirupsen/logrus"
)

func HasRight(account *models.Account, server *models.Server) bool {

	servers, err := repositories.FindAccountServers(account)
	if err != nil {
		log.Error(err)
		return false
	}
	log.Info(server.DiscordID)
	for _, accountServer := range servers {
		if server.DiscordID == accountServer.DiscordID {
			return true
		}
	}

	return false
}
