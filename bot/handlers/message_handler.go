package handlers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/cmd"
	"github.com/bwmarrin/discordgo"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/interfaces"
)

//MessageCreate is the handler called for every new message received by the bot on the guild
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check that the message is addressed to the bot
	if !strings.HasPrefix(strings.ToLower(m.Content), "assistant") {
		log.Info("commmand")
		return
	}

	params := strings.Split(m.Content, " ")
	if len(params) < 2 {
		return
	}

	genericCommand := cmd.GenericCommand{
		Session:     s,
		Message:     m,
		CommandType: params[1],
	}

	cmd, err := genericCommand.Build()
	log.Info("commmand: ", cmd)
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

	if err = interfaces.Command.Execute(cmd); err != nil {
		return
	}
}
