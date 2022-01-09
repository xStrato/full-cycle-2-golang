package commands

import (
	"github.com/asaskevich/govalidator"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

type CreateGenreCommand struct {
	*common.Command `json:"-" valid:"required"`
	Name            string              `json:"name" valid:"length(3|20)"`
	Categories      []entities.Category `json:"categories" valid:"length(1)"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCreateGenreCommand() *CreateGenreCommand {
	return &CreateGenreCommand{
		Command: common.NewCommand("CreateGenreCommand"),
	}
}
func NewCreateGenreCommandWithName(name string, categories []entities.Category) *CreateGenreCommand {
	return &CreateGenreCommand{
		Command: common.NewCommand("CreateGenreCommand"),
		Name:    name,
	}
}
func (c *CreateGenreCommand) IsValid() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}
