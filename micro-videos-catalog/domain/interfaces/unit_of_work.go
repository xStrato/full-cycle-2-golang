package interfaces

type UnitOfWork interface {
	Commit() error
}
