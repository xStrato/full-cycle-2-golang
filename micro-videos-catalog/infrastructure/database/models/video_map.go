package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id           string `gorm:"type:uuid;primaryKey"`
	Title        string
	Description  string
	YearLaunched int
	Opened       bool
	Rating       string
	Duration     float64
	Categories   []Category   `gorm:"foreignKey:Id"`
	Genres       []Genre      `gorm:"foreignKey:Id"`
	CastMembers  []CastMember `gorm:"foreignKey:Id"`
	VideoFiles   []VideoFile  `gorm:"foreignKey:Id"`
}
