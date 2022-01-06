package common

type GenericResult struct {
	Success bool
	Message string
	Data    interface{}
}

func NewGenericResult(success bool, message string, data interface{}) *GenericResult {
	return &GenericResult{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func (g *GenericResult) HasSuccess() bool {
	return g.Success
}

func (g *GenericResult) GetMessage() string {
	return g.Message
}

func (g *GenericResult) GetData() interface{} {
	return g.Data
}
