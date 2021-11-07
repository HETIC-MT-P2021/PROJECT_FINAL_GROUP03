package interfaces

type Command interface {
	Execute() error
}
