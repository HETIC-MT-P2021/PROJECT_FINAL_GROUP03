package services

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/security"
	"github.com/bwmarrin/discordgo"
)

// Sends link in a discord channel allowing front-end to identify user
func SendInterfaceLink(s *discordgo.Session, userID, channelID string) error {
	link := generateLink(userID)

	if _, err := s.ChannelMessageSend(channelID, link); err != nil {
		return errors.New("could not send link : " + err.Error())
	}

	return nil
}

func generateLink(authorID string) string {
	serverAdress := env.GetVariable("SERVER_ADDR_FRONT")
	authorIDHash := security.HashString(authorID)

	return serverAdress + "/login/" + authorIDHash
}
