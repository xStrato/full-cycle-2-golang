package commands

import (
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type CreateCategoryCommand struct {
	*common.Command
	Name string `valid:"required"`
}

func NewCreateCategoryCommand(name string) *CreateCategoryCommand {
	return &CreateCategoryCommand{
		Command: common.NewCommand("CreateCategoryCommand"),
		Name:    name,
	}
}
