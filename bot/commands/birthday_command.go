package commands

import (
	"errors"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/enum"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type BirthdayCommand struct {
	gc *GenericCommand
}

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
		"user_id": command.gc.Message.Author.ID,
		"server_id": command.gc.Message.GuildID,
		"channel_id": command.gc.Message.ChannelID,
		"birth_date": parsedDate.Format(time.RFC3339),
	}

	err = PerformRequest(enum.CreateUserBirthdayRoute, enum.Post, data)
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
