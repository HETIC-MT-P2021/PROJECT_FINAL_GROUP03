package models

type User struct {
	ID        uint64 `gorm:"primary_key"`
	DiscordID string `gorm:"uniqueIndex" json:"discord_id"`
	Name      string `json:"name"`

	/*Roles     []Role   `json:"roles"`
	Birthday  Birthday `json:"birthday"`*/
}
