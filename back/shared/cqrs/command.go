package cqrs

import (
	"fmt"
	"reflect"
)

// CommandMessage is the interface for all messages conveying in the command bus
type CommandMessage interface {
	CommandType() string
	Payload() interface{}
}

// CommandBus is the struct to represent the command bus, can only use commandHandlers
type CommandBus struct {
	handlers map[string]CommandHandler
}

// NewCommandBus is used to instanciate a new command bus
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

// Dispatch a command in the command bus, the handler corresponding to the command will be executed
func (b *CommandBus) Dispatch(command CommandMessage) (interface{}, error) {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}
	return nil, fmt.Errorf("the command bus does not have a handler for commands of type: %s", command.CommandType())
}

// RegisterHandler is used to add a handler to your command bus, it associates your handler with the command type provided
func (b *CommandBus) RegisterHandler(handler CommandHandler, command interface{}) error {
	typeName := reflect.TypeOf(command).Elem().Name()
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

// CommandDescriptor is the struct used to represent our commands
type CommandDescriptor struct {
	command interface{}
}

// NewCommandMessage takes an interface and returns a commandDescriptor that can be passed to the bus
func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

// CommandType returns the command type
func (c *CommandDescriptor) CommandType() string {
	return reflect.TypeOf(c.command).Elem().Name()
}

// Payload returns the actual command payload of the message.
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}
