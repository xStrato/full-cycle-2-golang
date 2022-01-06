package models

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id   string `gorm:"type:uuid;primaryKey"`
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		Id:   uuid.NewV4().String(),
		Name: name,
	}
}

func (e *Category) GetId() string {
	return e.Id
}

func (e *Category) SetId(id string) error {
	if len(id) <= 0 {
		return errors.New("'id' cannot be empty")
	}
	if !isValidUUID(id) {
		return errors.New("'id' is not a valid UUID")
	}
	e.Id = id
	return nil
}

func isValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err != nil
}
