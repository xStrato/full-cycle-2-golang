package handlers

import (
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common/interfaces"
)

type CategoryCommandHandler struct {
	repository interface{}
}

func NewCategoryCommandHandler(repo interface{}) *CategoryCommandHandler {
	return &CategoryCommandHandler{repo}
}

func (c *CategoryCommandHandler) Handle(cmd interfaces.Command) *common.GenericResult {
	if err := cmd.IsValid(); err != nil {
		return common.NewGenericResult(false, fmt.Sprintf("%v is not valid", cmd.GetCommandType()), err.Error())
	}
	return common.NewGenericResult(true, fmt.Sprintf("%v successfully executed", cmd.GetCommandType()), cmd)
}
