package models

type ForbiddenWordsListForm struct {
	ForbiddenWords string `json:"forbidden_words"`
	DiscordID      string `json:"discord_id"`
}
