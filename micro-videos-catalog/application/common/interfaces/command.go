package interfaces

type Command interface {
	GetCommandType() string
	IsValid() error
}
