package entities

import (
	"errors"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
)

type VideoFile struct {
	*common.Entity
	title    string
	url      string
	duration float64
}

func NewVideoFile(title string, duration float64) *VideoFile {
	return &VideoFile{
		Entity:   common.NewEntity(),
		title:    title,
		duration: duration,
	}
}

func NewVideoFileWithId(id, title string, duration float64) *VideoFile {
	return &VideoFile{
		Entity:   common.NewEntityWithId(id),
		title:    title,
		duration: duration,
	}
}
func NewVideoFileWithUrl(title, url string, duration float64) *VideoFile {
	return &VideoFile{
		Entity:   common.NewEntity(),
		title:    title,
		url:      url,
		duration: duration,
	}
}

func NewVideoFileWithIdAndUrl(id, title, url string, duration float64) *VideoFile {
	return &VideoFile{
		Entity:   common.NewEntityWithId(id),
		title:    title,
		url:      url,
		duration: duration,
	}
}

func (v *VideoFile) GetTitle() string {
	return v.title
}

func (v *VideoFile) SetTitle(title string) error {
	if len(title) <= 0 {
		return errors.New("'title' cannot be empty")
	}
	v.title = title
	return nil
}

func (v *VideoFile) GetUrl() string {
	return v.url
}

func (v *VideoFile) SetUrl(url string) {
	v.url = url
}

func (v *VideoFile) GetDuration() float64 {
	return v.duration
}

func (v *VideoFile) SetDuration(duration float64) {
	v.duration = duration
}
