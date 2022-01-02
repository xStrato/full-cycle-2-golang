package common

type GenericResult struct {
	success bool
	message string
	data    interface{}
}

func NewGenericResult(success bool, message string, data interface{}) *GenericResult {
	return &GenericResult{
		success: success,
		message: message,
		data:    data,
	}
}

func (g *GenericResult) HasSuccess() bool {
	return g.success
}

func (g *GenericResult) GetMessage() string {
	return g.message
}

func (g *GenericResult) GetData() interface{} {
	return g.data
}
