package services

import (
	"github.com/bwmarrin/discordgo"
)

func GetUserSession(token string) (*discordgo.Session, error) {
	return discordgo.New("Bearer " + token)
}