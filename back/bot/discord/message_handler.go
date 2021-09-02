package discord

import (
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/domain"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/services/accounts"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// MessageCreate is called when a new message is received by the bot
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	signUpifNotRegistered(m)

	if !strings.HasPrefix(strings.ToLower(m.Content), "/admin") {
		return
	}
	commandName := strings.Split(m.Content, " ")[1]

	var err error
	if commandName == "login" {
		channel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			log.Error(err)
			return
		}

		cmd := cqrs.NewCommandMessage(&domain.SendInterfaceLinkCommand{
			Session:   s,
			UserID:    m.Author.ID,
			ChannelID: channel.ID,
		})
		if _, err = infrastructure.CommandBus.Dispatch(cmd); err != nil {
			log.Error(err)
		}
	}

	if err != nil {
		log.Error(err)
		_, err := s.ChannelMessageSend(m.ChannelID, "Une erreur est survenue.")

		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		return
	}
}

func signUpifNotRegistered(m *discordgo.MessageCreate) {
	if !accounts.IsRegistered(m.Author.ID) {
		account := models.Account{
			Name:      m.Author.Username,
			DiscordID: m.Author.ID,
		}

		if err := repositories.PersistAccount(&account); err != nil {
			log.Error(err)
		}
	}
}
