package common

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Command struct {
	creationTime time.Time `valid:"-"`
	commandType  string    `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
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

func (c *Command) IsValid() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}
