package handlers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/commands"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	/*if !strings.HasPrefix(strings.ToLower(m.Content), "/admin") {
		// Admin Commands
	}*/
	params := strings.Split(m.Content, " ")

	genericCommand := commands.GenericCommand{
		Session:     s,
		Message:     m,
		CommandType: params[1],
	}

	cmd, err := genericCommand.Build()
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "Je n'ai pas réussi à trouver ce qu'il vous fallait.")
		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		return
	}
	if cmd == nil {
		return
	}
	if err = cmd.Execute(); err != nil {
		return
	}
}
