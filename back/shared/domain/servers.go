package domain

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	log "github.com/sirupsen/logrus"
)

type serverDomain struct{}

var Server serverDomain

// IsRegistered returns true if server with this discord id is found in database
func (sd serverDomain) IsRegistered(ID string) bool {
	res := models.Server{
		DiscordID: ID,
	}
	err := repositories.FindServerByDiscordID(&res)
	if err != nil {
		log.Error(err)
	}

	return res.Name != ""
}
