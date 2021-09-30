package domain

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/security"
	log "github.com/sirupsen/logrus"
)

type accountDomain struct{}

var Account accountDomain

// Returns true if user with this discord id is found in database
func (ad accountDomain) IsRegistered(ID string) bool {
	res := models.Account{
		DiscordID: ID,
	}
	err := repositories.FindAccountByDiscordID(&res)
	if err != nil {
		log.Error(err)
	}

	return res.Name != ""
}

// IsMember returns true if the user retrieved with the hash is a member of the server found with the discord id given
func (ad accountDomain) IsMember(hash, serverID string) bool {
	account := models.Account{
		DiscordID: hash,
	}
	if err := repositories.FindAccountByDiscordID(&account); err != nil {
		log.Error("user not found : ", err)
		return false
	}

	servers, err := repositories.FindAccountServers(&account)
	if err != nil {
		log.Error(err)
		return false
	}

	for _, server := range servers {
		if server.DiscordID == serverID {
			return true
		}
	}

	return false
}

func (ad accountDomain) GenerateLoginLink(authorID string) string {
	serverAdress := env.GetVariable("SERVER_ADDR_FRONT")
	authorIDHash := security.HashString(authorID)

	return serverAdress + "/login/" + authorIDHash
}
