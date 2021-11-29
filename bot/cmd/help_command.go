package cmd

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type HelpCommand struct {
	gc *GenericCommand
}

func (command HelpCommand) Execute() error {
	var err error
	message := discordgo.MessageEmbed{
		Type: "rich",
		Title: "Help",
		Description: "Each command must be prefixed with the `assistant ` string",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Name",
				Value:  "`help`",
				Inline: true,
			},
			{
				Name:   "Description",
				Value:  "Displays all available commands.",
				Inline: true,
			},
			{
				Name:   " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Value:  " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Inline: false,
			},
			{
				Name:   "Name",
				Value:  "`login`",
				Inline: true,
			},
			{
				Name:   "Description",
				Value:  "Go to website.",
				Inline: true,
			},
			{
				Name:   " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Value:  " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Inline: false,
			},
			{
				Name:   "Name",
				Value:  "`set-welcome_message`",
				Inline: true,
			},
			{
				Name:   "Description",
				Value:  "Change the welcome message, sent to every new user of the guild.",
				Inline: true,
			},
			{
				Name:   " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Value:  " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ",
				Inline: false,
			},
			{
				Name:   "Name",
				Value:  "`set-birthday`",
				Inline: true,
			},
			{
				Name:   "Description",
				Value:  "Set your birthday date in order for the bot to tell you Happy Birthday!",
				Inline: true,
			},
		},
	}
	if _, err = command.gc.Session.ChannelMessageSendEmbed(command.gc.Message.ChannelID, &message); err != nil {
		log.Error(err)
	}
	
	return err
}
