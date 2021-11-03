package commands

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type BirthdayCommand struct {
	gc *GenericCommand
}

func (command BirthdayCommand) Execute() error {
	params := strings.Split(command.gc.Message.Content, " ")
	log.Info(params)

	//_, err := command.gc.Session.ChannelMessageSendEmbed(command.gc.Message.ChannelID, &embed)

	return nil
}
