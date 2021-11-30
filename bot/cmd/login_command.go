package cmd

import (
	"errors"
	"github.com/JackMaarek/go-bot-utils/env"
)

type LoginCommand struct {
	gc *GenericCommand
}

func (command LoginCommand) Execute() error {
	var frontUrl string
	if frontUrl = env.GetVariable("FRONT_URL"); frontUrl == "" {
		return errors.New("MISSING ENV VARIABLE FRONT_URL")
	}

	if _, err := command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, frontUrl); err != nil {
		return err
	}

	return nil
}
