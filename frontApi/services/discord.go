package services

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func GetUserSession(c *gin.Context) (*discordgo.Session, error) {
	accessToken := c.GetHeader("Authorization")

	if "" == accessToken {
		return &discordgo.Session{}, errors.New("authorizaion code needed")
	}

	return discordgo.New("Bearer " + accessToken)
}

func GetUserGuildByID(session *discordgo.Session, guildID string) (*discordgo.UserGuild, error) {
	var userGuild *discordgo.UserGuild
	userGuilds, err := session.UserGuilds(100, "", "")
	if err != nil {
		return userGuild, err
	}

	for _, guild := range userGuilds {
		if guild.ID == guildID {
			userGuild = guild
			break
		}
	}

	return userGuild, nil
}