package services

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

func MemberHasPermission(s *discordgo.Session, guildID string, permission int64) (bool, error) {
	user := GetUser(s)

	member, err := s.State.Member(guildID, user.ID)
	if err != nil {
		if member, err = s.GuildMember(guildID, user.ID); err != nil {
			return false, err
		}
	}

	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}

	return false, nil
}

func GetUser(session *discordgo.Session) *discordgo.User {
	user, err := session.User("@me")
	if err != nil {
		log.Error("error while fetching user", err)
	}
	return user
}
