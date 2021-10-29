package models

type Server struct {
	DiscordID 		string `gorm:"primary_key"`
	WelcomeMessage 	string
}
