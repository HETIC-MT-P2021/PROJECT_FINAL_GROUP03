package enum

type Routes string

const (
	GetServerResource          = "/api/v1/servers"
	ChangeWelcomeMessageRoute  = "/api/v1/commands/change-welcome-message"
	ChangeBirthdayMessageRoute = "/api/v1/commands/change-birthday-message"
	CreateUserBirthdayRoute    = "/api/v1/birthdays"
	RolesResourceRoute         = "/api/v1/roles"
)

type Methods string

const (
	Get    = "GET"
	Post   = "POST"
	Patch  = "PATCH"
	Delete = "DELETE"
)
