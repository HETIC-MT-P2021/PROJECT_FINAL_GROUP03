package cmd

import (
	"errors"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/enum"
	"github.com/JackMaarek/go-bot-utils/helpers"
)

// BirthdayCommand holds needed informations to change the message author's birth date
type BirthdayCommand struct {
	gc *GenericCommand
}

// Execute the birthday command to set birth date for message author
func (command BirthdayCommand) Execute() error {
	params := strings.Split(command.gc.Message.Content, " ")

	if len(params) < 3 {
		return errors.New("not enough arguments")
	}

	parsedDate, err := time.Parse("02-01-2006", strings.Replace(params[2], "/", "-", 2))
	if err != nil {
		log.Info(err)
		return errors.New("could not parse birth date")
	}

	data := map[string]string{
		"user_id":    command.gc.Message.Author.ID,
		"server_id":  command.gc.Message.GuildID,
		"channel_id": command.gc.Message.ChannelID,
		"birth_date": parsedDate.Format(time.RFC3339),
	}

	_, err = helpers.PerformRequest(enum.CreateUserBirthdayRoute, enum.Post, data)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	_, err = command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, "You will be celebrated bitch")
	if err != nil {
		log.Info(err.Error())
		return err
	}

	return nil
}
