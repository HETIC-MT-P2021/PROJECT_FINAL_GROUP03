package interfaces

import "github.com/bwmarrin/discordgo"

type Discord interface {
	ChannelMessageSend(string, string) (*discordgo.Message, error)
	ChannelMessageSendEmbed(string, *discordgo.MessageEmbed) (*discordgo.Message, error)
}
