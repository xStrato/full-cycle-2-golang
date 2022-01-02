package commands

import (
	"github.com/asaskevich/govalidator"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type CreateCategoryCommand struct {
	*common.Command `valid:"required"`
	Name            string `valid:"required,length(3|20)"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCreateCategoryCommand(name string) *CreateCategoryCommand {
	return &CreateCategoryCommand{
		Command: common.NewCommand("CreateCategoryCommand"),
		Name:    name,
	}
}

func (c *CreateCategoryCommand) IsValid() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}
