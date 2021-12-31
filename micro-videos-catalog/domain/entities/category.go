package entities

import "errors"

type Category struct {
	id, name string
}

func NewCategory(name string) *Category {
	return &Category{name: name}
}

func NewFullCategory(id, name string) *Category {
	return &Category{id, name}
}

func (c *Category) GetId() string {
	return c.id
}

func (c *Category) SetId(id string) error {
	if len(id) <= 0 {
		return errors.New("'id' cannot be empty")
	}
	c.id = id
	return nil
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
