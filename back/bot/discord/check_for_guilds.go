package discord

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/services/servers"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func checkForGuilds(s *discordgo.Session) {
	// Get guilds
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		log.Error(err)
		return
	}

	// For each guild, check if in db and save it if it is not
	for _, guild := range guilds {
		server := models.Server{
			Name:      guild.Name,
			DiscordID: guild.ID,
		}
		log.Info("guild", server.Name)
		if servers.IsRegistered(server.DiscordID) {
			log.Info("already exists : ", server.DiscordID)
			continue
		}
		if err := repositories.PersistServer(&server); err != nil {
			log.Error(err)
		}
	}
}
