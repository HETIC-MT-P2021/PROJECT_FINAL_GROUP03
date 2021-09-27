package domain

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/services"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/bwmarrin/discordgo"
)

type ChangeWelcomeMessageCommand struct {
	ServerDiscordID,
	WelcomeMessage string
	Session *discordgo.Session
}

type ChangeWelcomeMessageCommandHandler struct{}

func (handler *ChangeWelcomeMessageCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *ChangeWelcomeMessageCommand:
		err := services.ChangeWelcomeMessage(cmd.ServerDiscordID, cmd.WelcomeMessage, cmd.Session)
		return nil, err
	default:
		return nil, nil
	}
}
