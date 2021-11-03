package commands

import (
	"errors"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/interfaces"
	"github.com/bwmarrin/discordgo"
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
	}
	return nil, errors.New("Could not find command")
}