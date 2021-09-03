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
	signUpifNotRegistered(s)

	if !strings.HasPrefix(strings.ToLower(m.Content), "/admin") {
		return
	}
	params := strings.Split(m.Content, " ")
	if len(params) < 2 {
		return
	}
	commandName := params[1]

	switch commandName {
	case "login":
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

		if err != nil {
			log.Error(err)
			_, err := s.ChannelMessageSend(m.ChannelID, "Une erreur est survenue.")

			if err != nil {
				log.Error("sendMessageErr: ", err)
			}
			return
		}
		break
	case "set-welcome-message":
		params := params[2:]
		message := strings.Join(params, " ")
		cmd := cqrs.NewCommandMessage(&domain.ChangeWelcomeMessageCommand{
			Session:   				s,
			ServerDiscordID:  m.GuildID,
			WelcomeMessage: 	message,
		})
		
		_, err := infrastructure.CommandBus.Dispatch(cmd)
		if err != nil {
			log.Error(err)
			_, err := s.ChannelMessageSend(m.ChannelID, "Une erreur est survenue.")

			if err != nil {
				log.Error("sendMessageErr: ", err)
			}
		} else {
			_, err := s.ChannelMessageSend(m.ChannelID, "Nouveau message sauvegardé")

			if err != nil {
				log.Error("sendMessageErr: ", err)
			}
		}
		break
	}
}

func signUpifNotRegistered(s *discordgo.Session) {
	if !accounts.IsRegistered(s.State.User.ID) {
		account := models.Account{
			Name:      s.State.User.Username,
			DiscordID: s.State.User.ID,
		}

		if err := repositories.PersistAccount(&account); err != nil {
			log.Error(err)
		}
	}
}
