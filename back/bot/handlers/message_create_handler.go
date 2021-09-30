package handlers

import (
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/domain"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/security"
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

		cmd := cqrs.NewCommandMessage(&commands.SendInterfaceLinkCommand{
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
	case "set-welcome-message":
		params := params[2:]
		message := strings.Join(params, " ")
		cmd := cqrs.NewCommandMessage(&commands.ChangeWelcomeMessageCommand{
			Session:         s,
			ServerDiscordID: m.GuildID,
			WelcomeMessage:  message,
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
	}
}

func signUpifNotRegistered(m *discordgo.MessageCreate) {
	if !domain.Account.IsRegistered(security.HashString(m.Author.ID)) {
		account := models.Account{
			Name:      m.Author.Username,
			DiscordID: security.HashString(m.Author.ID),
		}

		if err := repositories.PersistAccount(&account); err != nil {
			log.Error(err)
			return
		}
	}

	account := models.Account{
		DiscordID: security.HashString(m.Author.ID),
	}
	if err := repositories.FindAccountByDiscordID(&account); err != nil {
		log.Error("account not found : ", err)
		return
	}
	server := models.Server{
		DiscordID: m.GuildID,
	}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		log.Error("server not found", err)
		return
	}
	account.Servers = append(account.Servers, &server)
	if err := repositories.UpdateAccount(&account); err != nil {
		log.Error("couldn't update account servers", err)
	}
}
