package commands

import (
	"github.com/asaskevich/govalidator"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type CreateCategoryCommand struct {
	*common.Command `json:"-" valid:"required"`
	Name            string `json:"name" valid:"required,length(3|20)"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCreateCategoryCommand() *CreateCategoryCommand {
	return &CreateCategoryCommand{
		Command: common.NewCommand("CreateCategoryCommand"),
	}
}
func NewCreateCategoryCommandWithName(name string) *CreateCategoryCommand {
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
