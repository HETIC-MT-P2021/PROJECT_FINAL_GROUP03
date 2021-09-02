package models

type Server struct {
	ID							uint64			`gorm:"primary_key"`
	DiscordID,
	WelcomeMessage,
	Name 						string

	Accounts 				[]*Account 	`gorm:"many2many:server_accounts;"`
}