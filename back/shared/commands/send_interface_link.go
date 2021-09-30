package commands

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/services"
	"github.com/bwmarrin/discordgo"
)

type SendInterfaceLinkCommand struct {
	UserID,
	ChannelID string
	Session *discordgo.Session
}

type SendInterfaceLinkCommandHandler struct{}

func (handler *SendInterfaceLinkCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *SendInterfaceLinkCommand:
		err := services.SendInterfaceLink(cmd.Session, cmd.UserID, cmd.ChannelID)
		return nil, err
	default:
		return nil, nil
	}
}
