package services

import (
	"errors"
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/JackMaarek/go-bot-utils/models"
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
			if birthday.BirthDate.Before(time.Now()) && !birthday.MessageSent {
				if _, err := session.ChannelMessageSend(birthday.ChannelID, server.BirthdayMessage); err != nil {
					log.Info(fmt.Sprintf("Could not send birthday message to user id %s", birthday.UserID))
					continue
				} else {
					birthday.MessageSent = true
					if err := repositories.UpdateUserBirthday(&birthday); err != nil {
						log.Info(fmt.Sprintf("Could not update user's birthday with user id %s", birthday.UserID))
					}
				}
			}
		}
	}
	return nil
}
