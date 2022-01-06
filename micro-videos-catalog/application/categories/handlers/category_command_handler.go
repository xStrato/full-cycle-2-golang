package handlers

import (
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
	app_interfaces "github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common/interfaces"
	interfaces "github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/models"
)

type CategoryCommandHandler struct {
	repository interfaces.CategoryRepository
}

func NewCategoryCommandHandler(repo interfaces.CategoryRepository) *CategoryCommandHandler {
	return &CategoryCommandHandler{repo}
}

func (c *CategoryCommandHandler) Handle(cmd app_interfaces.Command) *common.GenericResult {
	switch command := cmd.(type) {
	case *commands.CreateCategoryCommand:
		return c.handleCreateCommand(command)
	}
	return common.NewGenericResult(false, fmt.Sprintf("'%v' is not supported", cmd.GetCommandType()), nil)
}

func (c *CategoryCommandHandler) handleCreateCommand(cmd *commands.CreateCategoryCommand) *common.GenericResult {
	if err := cmd.IsValid(); err != nil {
		return common.NewGenericResult(false, fmt.Sprintf("%v state is invalid", cmd.GetCommandType()), err.Error())
	}
	model := models.NewCategory(cmd.Name)
	c.repository.Add(model)
	if err := c.repository.Commit(); err != nil {
		return common.NewGenericResult(false, fmt.Sprintf("%v cannot be fully executed", cmd.GetCommandType()), err.Error())
	}
	return common.NewGenericResult(true, fmt.Sprintf("%v was successfully executed", cmd.GetCommandType()), model)
}
