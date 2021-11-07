package handlers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/JackMaarek/go-bot-utils/models"
)

func ForbiddenMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// avoid moderating bot messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// get words list for server
	var server models.Server
	server.DiscordID = m.GuildID
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		return
	}

	forbiddenWords := strings.Split(server.ForbiddenWords, ",")

	for _, word := range forbiddenWords {
		if strings.Contains(m.Content, word) {
			if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
				log.Error("moderation issue on message : ", m.ID, err)
			}

			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" veuillez vous calmer, paix et amour ❤")

			return
		}
	}
}
