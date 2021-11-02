package services

import (
	"errors"
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"time"
)

func RemindBirthdays(session *discordgo.Session) error {
	var servers []models.Server
	var birthdays []models.Birthday
	if err := repositories.FindAllServers(&servers); err != nil {
		return errors.New("could not find servers")
	}

	if err := repositories.FindAllBirthdays(&birthdays); err != nil {
		return errors.New("could not find birthdays")
	}

	for _, server := range servers {
		if server.BirthdayMessage == "" {
			log.Warn(fmt.Sprintf("No birthday message set for server id %s", server.DiscordID))
			continue
		}
		for _, birthday := range birthdays {
			if birthday.BirthDate.Before(time.Now()) && false == birthday.MessageSent {
				if _, err := session.ChannelMessageSend(birthday.ChannelID, server.BirthdayMessage); err != nil {
					return errors.New(fmt.Sprintf("Could not send birthday message to user id %s", birthday.UserID))
				}
			}
		}
	}
	return nil
}
