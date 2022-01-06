package interfaces

import (
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type RequestCommandHandler interface {
	Handle(cmd Command) *common.GenericResult
}
