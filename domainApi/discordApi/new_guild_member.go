package discordApi

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func GuildMemberAdd(s *discordgo.Session, data *discordgo.GuildMemberAdd) {
	// Get welcome message and send message with it
	server := models.Server{
		DiscordID: data.GuildID,
	}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		log.Error(err)
		return
	}

	channel, err := s.UserChannelCreate(data.User.ID)
	if err != nil {
		log.Error(err)
		return
	}
	if _, err := s.ChannelMessageSend(channel.ID, server.WelcomeMessage); err != nil {
		log.Error(err)
	}
}