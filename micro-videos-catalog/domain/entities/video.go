package entities

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces"
)

type Video struct {
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
func (v *Video) SetCategories(c []Category) error {
	if c == nil || len(c) <= 0 {
		return errors.New("'categories' cannot be empty or nil")
	}
	v.categories = c
	return nil
}
func (v *Video) GetGenres() []Genre {
	return v.genres
}
func (v *Video) SetGenres(g []Genre) error {
	if g == nil || len(g) <= 0 {
		return errors.New("'genres' cannot be empty or nil")
	}
	v.genres = g
	return nil
}
func (v *Video) GetCastMembers() []CastMember {
	return v.castMembers
}
func (v *Video) SetCastMembers(c []CastMember) error {
	if c == nil || len(c) <= 0 {
		return errors.New("'castMembers' cannot be empty or nil")
	}
	v.castMembers = c
	return nil
}

func (v *Video) AddT(t interfaces.Entity) error {
	if reflect.ValueOf(t).IsNil() {
		return fmt.Errorf("cannot Add 'nil' value")
	}
	switch entity := t.(type) {
	case *Category:
		v.addCategory(entity)
		return nil
	case *Genre:
		v.addGenre(entity)
		return nil
	case *CastMember:
		v.addCastMember(entity)
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
		v.removeCategory(t.GetId())
		return nil
	case *Genre:
		v.removeGenre(t.GetId())
		return nil
	case *CastMember:
		v.removeCastMember(t.GetId())
		return nil
	}
	return fmt.Errorf("type '%v' is not supported", t)
}

func (v *Video) removeCategory(id string) {
	for i := range v.categories {
		if v.categories[i].GetId() == id {
			v.categories = append(v.categories[:i], v.categories[i+1:]...)
			return
		}
	}
}
func (v *Video) removeGenre(id string) {
	for i := range v.genres {
		if v.genres[i].GetId() == id {
			v.genres = append(v.genres[:i], v.genres[i+1:]...)
			return
		}
	}
}
func (v *Video) removeCastMember(id string) {
	for i := range v.castMembers {
		if v.castMembers[i].GetId() == id {
			v.castMembers = append(v.castMembers[:i], v.castMembers[i+1:]...)
			return
		}
	}
}

func (v *Video) addCategory(c *Category) {
	for _, value := range v.categories {
		if value.GetId() == c.GetId() {
			return
		}
	}
	v.categories = append(v.categories, *c)
}

func (v *Video) addGenre(g *Genre) {
	for _, value := range v.genres {
		if value.GetId() == g.GetId() {
			return
		}
	}
	v.genres = append(v.genres, *g)
}

func (v *Video) addCastMember(c *CastMember) {
	for _, value := range v.castMembers {
		if value.GetId() == c.GetId() {
			return
		}
	}
	v.castMembers = append(v.castMembers, *c)
}
