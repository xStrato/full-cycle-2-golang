package interfaces

type Repository interface {
	UnitOfWork
	GetAll() ([]Entity, error)
	GetById(id string) (Entity, error)
	Add(e Entity) error
	Update(e Entity) (Entity, error)
	Delete(id string) error
}
