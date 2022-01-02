package entities

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces"
	"gorm.io/gorm"
)

type Video struct {
	*gorm.Model
	*common.Entity
	title        string
	description  string
	yearLaunched int
	opened       bool
	rating       string
	duration     float64
	categories   []Category
	genres       []Genre
	castMembers  []CastMember
	videoFiles   []VideoFile
}

func NewVideo(title, description string, year int, opened bool) *Video {
	return &Video{
		Entity:       common.NewEntity(),
		title:        title,
		description:  description,
		yearLaunched: year,
		opened:       opened,
		categories:   make([]Category, 0),
	}
}

func NewVideoWithId(id, title, description string, year int, opened bool) *Video {
	return &Video{
		Entity:       common.NewEntityWithId(id),
		title:        title,
		description:  description,
		yearLaunched: year,
		opened:       opened,
		categories:   make([]Category, 0),
	}
}

func NewVideoWithDuration(title, description string, year int, opened bool, duration float64) *Video {
	return &Video{
		Entity:       common.NewEntity(),
		title:        title,
		description:  description,
		yearLaunched: year,
		opened:       opened,
		duration:     duration,
		categories:   make([]Category, 0),
	}
}

func NewVideoComplete(title, description string, year int, opened bool, rating string, duration float64) *Video {
	return &Video{
		Entity:       common.NewEntity(),
		title:        title,
		description:  description,
		yearLaunched: year,
		opened:       opened,
		rating:       rating,
		duration:     duration,
		categories:   make([]Category, 0),
	}
}

func NewVideoCompleteWithId(id, title, description string, year int, opened bool, rating string, duration float64) *Video {
	return &Video{
		Entity:       common.NewEntityWithId(id),
		title:        title,
		description:  description,
		yearLaunched: year,
		opened:       opened,
		rating:       rating,
		duration:     duration,
		categories:   make([]Category, 0),
	}
}

func (v *Video) GetTitle() string {
	return v.title
}

func (v *Video) SetTitle(title string) error {
	if len(title) <= 0 {
		return errors.New("'title' cannot be empty")
	}
	v.title = title
	return nil
}

func (v *Video) GetYearLaunched() int {
	return v.yearLaunched
}

func (v *Video) SetYearLaunched(yearLaunched int) error {
	if yearLaunched >= 0 && yearLaunched <= 1700 {
		return errors.New("'yearLaunched' must be greater than 1700")
	}
	if yearLaunched > time.Now().Year() {
		return errors.New("'yearLaunched' is greater than current year")
	}
	v.yearLaunched = yearLaunched
	return nil
}

func (v *Video) GetOpened() bool {
	return v.opened
}

func (v *Video) SetOpened(opened bool) {
	v.opened = opened
}

func (v *Video) GetRating() string {
	return v.rating
}

func (v *Video) SetRating(rating string) {
	v.rating = rating
}

func (v *Video) GetDescription() string {
	return v.description
}

func (v *Video) SetDescription(description string) {
	v.description = description
}

func (v *Video) GetDuration() float64 {
	return v.duration
}

func (v *Video) SetDuration(duration float64) {
	v.duration = duration
}

func (v *Video) GetCategories() []Category {
	return v.categories
}
func (v *Video) SetCategories(c *[]Category) error {
	if c == nil || len(*c) <= 0 {
		return errors.New("'categories' cannot be empty or nil")
	}
	v.categories = *c
	return nil
}
func (v *Video) GetGenres() []Genre {
	return v.genres
}
func (v *Video) SetGenres(g *[]Genre) error {
	if g == nil || len(*g) <= 0 {
		return errors.New("'genres' cannot be empty or nil")
	}
	v.genres = *g
	return nil
}

func (v *Video) GetCastMembers() []CastMember {
	return v.castMembers
}
func (v *Video) SetCastMembers(c *[]CastMember) error {
	if c == nil || len(*c) <= 0 {
		return errors.New("'castMembers' cannot be empty or nil")
	}
	v.castMembers = *c
	return nil
}

func (v *Video) AddT(t interfaces.Entity) error {
	if reflect.ValueOf(t).IsNil() {
		return fmt.Errorf("cannot Add 'nil' value")
	}
	switch entity := t.(type) {
	case *Category:
		query := linq.From(v.categories)
		if _, ok := contains(t.GetId(), &query); !ok {
			v.categories = append(v.categories, *entity)
		}
		return nil
	case *Genre:
		query := linq.From(v.genres)
		if _, ok := contains(t.GetId(), &query); !ok {
			v.genres = append(v.genres, *entity)
		}
		return nil
	case *CastMember:
		query := linq.From(v.castMembers)
		if _, ok := contains(t.GetId(), &query); !ok {
			v.castMembers = append(v.castMembers, *entity)
		}
		return nil
	}
	return fmt.Errorf("type '%v' is not supported", t)
}

func (v *Video) RemoveT(t interfaces.Entity) error {
	if reflect.ValueOf(t).IsNil() {
		return fmt.Errorf("cannot Remove 'nil' value")
	}
	switch t.(type) {
	case *Category:
		query := linq.From(v.categories)
		if i, ok := contains(t.GetId(), &query); ok {
			v.categories = append(v.categories[:i], v.categories[i+1:]...)
		}
		return nil
	case *Genre:
		query := linq.From(v.genres)
		if i, ok := contains(t.GetId(), &query); ok {
			v.genres = append(v.genres[:i], v.genres[i+1:]...)
		}
		return nil
	case *CastMember:
		query := linq.From(v.castMembers)
		if i, ok := contains(t.GetId(), &query); ok {
			v.castMembers = append(v.castMembers[:i], v.castMembers[i+1:]...)
		}
		return nil
	}
	return fmt.Errorf("type '%v' is not supported", t)
}

func contains(id string, s *linq.Query) (int, bool) {
	index := s.IndexOfT(func(e interfaces.Entity) bool {
		return e.GetId() == id
	})
	return index, index >= 0
}
