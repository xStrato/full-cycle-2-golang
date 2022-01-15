package repositories

import (
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/contexts"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/models"
)

type CategoryRepository struct {
	dbContext *contexts.VideoCatalogContext
}

func NewCategoryRepository(ctx *contexts.VideoCatalogContext) *CategoryRepository {
	return &CategoryRepository{ctx}
}

func (cr *CategoryRepository) GetAll() ([]interfaces.Entity, error) {
	return make([]interfaces.Entity, 0), nil
}
func (cr *CategoryRepository) GetById(id string) (interfaces.Entity, error) {
	return *entities.NewCategory("teste"), nil
}
func (cr *CategoryRepository) Add(e interfaces.Entity) error {
	model, ok := e.(*models.Category)
	if !ok {
		return fmt.Errorf("cannot cast '%v' as Category entity", e)
	}
	db := cr.dbContext.GetDBConnection()
	if err := db.Create(&model).Error; err != nil {
		return err
	}
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
