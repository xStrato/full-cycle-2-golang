package repositories

import (
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces"
)

type CategoryRepository struct {
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (cr *CategoryRepository) GetAll() ([]interfaces.Entity, error) {
	return make([]interfaces.Entity, 0), nil
}
func (cr *CategoryRepository) GetById(id string) (interfaces.Entity, error) {
	return *entities.NewCategory("teste"), nil
}
func (cr *CategoryRepository) Add(e interfaces.Entity) error {
	category, ok := e.(entities.Category)
	if !ok {
		// log.Fatalf("CategoryRepository Add, Error")
	}
	fmt.Println(category)
	return nil
}
func (cr *CategoryRepository) Update(e interfaces.Entity) (interfaces.Entity, error) {
	return e, nil
}
func (cr *CategoryRepository) Delete(id string) error {
	return nil
}

func (cr *CategoryRepository) Commit() error {
	return nil
}
