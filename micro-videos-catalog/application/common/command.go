package common

import (
	"time"
)

type Command struct {
	creationTime time.Time
	commandType  string
}

func NewCommand(name string) *Command {
	return &Command{time.Now(), name}
}

func (c *Command) GetCreationTime() time.Time {
	return c.creationTime
}

func (c *Command) GetCommandType() string {
	return c.commandType
}
