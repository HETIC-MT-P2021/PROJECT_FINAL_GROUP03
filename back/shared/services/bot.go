package services

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/domain"
	"github.com/bwmarrin/discordgo"
)

// Sends link in a discord channel allowing front-end to identify user
func SendInterfaceLink(s *discordgo.Session, userID, channelID string) error {
	link := domain.Account.GenerateLoginLink(userID)

	if _, err := s.ChannelMessageSend(channelID, link); err != nil {
		return errors.New("could not send link : " + err.Error())
	}

	return nil
}
