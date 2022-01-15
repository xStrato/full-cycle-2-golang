package models

import (
	"gorm.io/gorm"
)

type VideoFile struct {
	gorm.Model
	Id       string `gorm:"type:uuid;primaryKey"`
	Title    string
	Url      string
	Duration float64
}
