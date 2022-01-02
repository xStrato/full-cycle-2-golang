package common

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Entity struct {
	id string `gorm:"primaryKey"`
}

func NewEntity() *Entity {
	return &Entity{id: uuid.NewV4().String()}
}

func NewEntityWithId(id string) *Entity {
	return &Entity{id: id}
}

func (e *Entity) GetId() string {
	return e.id
}

func (e *Entity) SetId(id string) error {
	if len(id) <= 0 {
		return errors.New("'id' cannot be empty")
	}
	if !isValidUUID(id) {
		return errors.New("'id' is not a valid UUID")
	}
	e.id = id
	return nil
}

func isValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err != nil
}
