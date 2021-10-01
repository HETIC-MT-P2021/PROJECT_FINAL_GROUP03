package models

type Server struct {
	ID uint64 `gorm:"primary_key" json:"id"`
	DiscordID string `json:"discord_id"`
	WelcomeMessage string `json:"welcome_message"`
	Name string			`json:"name"`

	Accounts []*Account `gorm:"many2many:server_accounts;"`
}
