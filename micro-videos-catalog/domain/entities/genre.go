package entities

import (
	"errors"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
)

type Genre struct {
	*common.Entity
	name       string
	categories []Category
}

func NewGenre(name string) *Genre {
	return &Genre{
		Entity:     common.NewEntity(),
		name:       name,
		categories: make([]Category, 0),
	}
}

func NewGenreWithId(id, name string) *Genre {
	return &Genre{
		Entity:     common.NewEntityWithId(id),
		name:       name,
		categories: make([]Category, 0),
	}
}
func NewGenreWithCategories(name string, c []Category) *Genre {
	return &Genre{
		Entity:     common.NewEntity(),
		name:       name,
		categories: c,
	}
}

func NewGenreWithIdAndCategories(id, name string, c []Category) *Genre {
	return &Genre{
		Entity:     common.NewEntityWithId(id),
		name:       name,
		categories: c,
	}
}

func (g *Genre) GetName() string {
	return g.name
}

func (g *Genre) SetName(n string) error {
	if len(n) <= 0 {
		return errors.New("'name' cannot be empty")
	}
	g.name = n
	return nil
}

func (g *Genre) GetCategories() []Category {
	return g.categories
}

func (g *Genre) SetCategories(c []Category) error {
	if c == nil || len(c) <= 0 {
		return errors.New("'categories' cannot be empty or nil")
	}
	g.categories = c
	return nil
}

func (g *Genre) AddCategory(c *Category) error {
	if c == nil {
		return errors.New("'category' cannot be nil")
	}
	g.categories = append(g.categories, *c)
	return nil
}

func (g *Genre) RemoveCategory(c *Category) error {
	if c == nil {
		return errors.New("'category' cannot be nil")
	}
	for i := range g.categories {
		if g.categories[i].GetId() == c.GetId() {
			g.categories = append(g.categories[:i], g.categories[i+1:]...)
			return nil
		}
	}
	return nil
}
