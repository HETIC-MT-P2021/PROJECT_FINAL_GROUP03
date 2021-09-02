package models

type Account struct {
	ID 				uint64
	Name,
	DiscordID string

	Servers 	[]*Server 	`gorm:"many2many:server_accounts;"`
}