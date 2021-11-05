package models

type ChangeBirthdayMessage struct {
	BirthdayMessage string `json:"birthday_message"`
	DiscordID       string `json:"discord_id"`
}
