package models

type Server struct {
	DiscordID       string `json:"discord_id"`
	Name            string `json:"name"`
	WelcomeMessage  string `json:"welcome_message"`
	BirthdayMessage string `json:"birthday_message"`
	ForbiddenWords  string `json:"forbidden_words"`
}
