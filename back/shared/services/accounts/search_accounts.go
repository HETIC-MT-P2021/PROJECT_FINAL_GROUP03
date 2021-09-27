package accounts

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	log "github.com/sirupsen/logrus"
)

// Returns true if user with this discord id is found in database
func IsRegistered(ID string) bool {
	res := models.Account{
		DiscordID: ID,
	}
	err := repositories.FindAccountByDiscordID(&res)
	if err != nil {
		log.Error(err)
	}

	return res.Name != ""
}