package handlers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/models"
)

// CheckForGuilds updates bot guilds list
func CheckForGuilds(s *discordgo.Session) {
	// Get guilds
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		log.Error(err)
		return
	}

	// For each guild, check if in db and save it if it is not
	for _, guild := range guilds {
		server := models.Server{
			DiscordID: guild.ID,
		}

		if err := repositories.FindServerByDiscordID(&server); err == nil {
			continue
		}
		if err := repositories.PersistServer(&server); err != nil {
			log.Error(err)
		}
	}
}
