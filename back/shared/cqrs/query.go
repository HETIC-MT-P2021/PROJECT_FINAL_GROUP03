package cqrs

import (
	"fmt"
	"reflect"
)

// QueryMessage is the interface for all messages conveying in the command bus
type QueryMessage interface {
	Payload() interface{}
	QueryType() string
}

// QueryBus is the struct to represent the command bus, can only use commandHandlers
type QueryBus struct {
	handlers map[string]QueryHandler
}

// NewQueryBus is used to instanciate a new command bus
func NewQueryBus() *QueryBus {
	cBus := &QueryBus{
		handlers: make(map[string]QueryHandler),
	}

	return cBus
}

// Dispatch a command in the command bus, the handler corresponding to the command will be executed
func (b *QueryBus) Dispatch(query QueryMessage) (interface{}, error) {
	if handler, ok := b.handlers[query.QueryType()]; ok {
		return handler.Handle(query)
	}
	return nil, fmt.Errorf("the query bus does not have a handler for query of type: %s", query.QueryType())
}

// RegisterHandler is used to add a handler to your command bus, it associates your handler with the command type provided
func (b *QueryBus) RegisterHandler(handler QueryHandler, query interface{}) error {
	typeName := reflect.TypeOf(query).Elem().Name()
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate query handler registration with query bus for query of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

// QueryDescriptor is the struct used to represent our commands
type QueryDescriptor struct {
	query interface{}
}

// NewQueryMessage takes an interface and returns a commandDescriptor that can be passed to the bus
func NewQueryMessage(query interface{}) *QueryDescriptor {
	return &QueryDescriptor{
		query: query,
	}
}

// QueryType returns the command type
func (c *QueryDescriptor) QueryType() string {
	return reflect.TypeOf(c.query).Elem().Name()
}

// Payload returns the actual command payload of the message.
func (c *QueryDescriptor) Payload() interface{} {
	return c.query
}
