package cmd

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/JackMaarek/go-bot-utils/interfaces"
)

type GenericCommand struct {
	Session     interfaces.Discord
	Message     *discordgo.MessageCreate
	CommandType string
}

// Build a command depending on the analysis result we give
func (gc *GenericCommand) Build() (interfaces.Command, error) {
	switch gc.CommandType {

	case "set-birthday":
		return BirthdayCommand{gc: gc}, nil

	case "set-welcome_message":
		return WelcomeCommand{gc: gc}, nil

	case "help":
		return HelpCommand{gc: gc}, nil

	case "login":
		return LoginCommand{gc: gc}, nil
	}

	return nil, errors.New("Could not find command")
}
