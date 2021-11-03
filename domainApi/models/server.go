package models

type Server struct {
	DiscordID 		string `gorm:"primary_key" json:"discord_id"`
	WelcomeMessage 	string `json:"welcome_message"`
}
