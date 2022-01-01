package interfaces

type Entity interface {
	GetId() string
	SetId(id string) error
}
