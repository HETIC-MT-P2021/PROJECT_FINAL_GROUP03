package cqrs

// CommandHandler is the interface to give to the bus in order to register a command&handler
type CommandHandler interface {
	Handle(CommandMessage) (interface{}, error)
}

// QueryHandler is the interface to give to the bus in order to register a query&handler
type QueryHandler interface {
	Handle(QueryMessage) (interface{}, error)
}
