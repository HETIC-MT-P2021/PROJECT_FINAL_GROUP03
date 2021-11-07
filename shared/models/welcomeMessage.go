package models

type ChangeWelcomeMessageForm struct {
	WelcomeMessage string `json:"welcome_message"`
	DiscordID      string `json:"discord_id"`
}
