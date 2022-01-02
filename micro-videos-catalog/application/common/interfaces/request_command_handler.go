package interfaces

import (
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type RequestCommandHandler interface {
	Handle(c Command) common.GenericResult
}
