package discord

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func GuildMemberAdd(s *discordgo.Session, data *discordgo.GuildMemberAdd) {
	log.Info("new member", data)
}
