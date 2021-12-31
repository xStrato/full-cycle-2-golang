package entities

import (
	"errors"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
)

type Category struct {
	*common.Entity
	name string
}

func NewCategory(name string) *Category {
	return &Category{
		Entity: common.NewEntity(),
		name:   name,
	}
}

func NewCategoryWithId(id, name string) *Category {
	return &Category{
		Entity: common.NewEntityWithId(id),
		name:   name,
	}
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) SetName(n string) error {
	if len(n) <= 0 {
		return errors.New("'name' cannot be empty")
	}
	c.name = n
	return nil
}
