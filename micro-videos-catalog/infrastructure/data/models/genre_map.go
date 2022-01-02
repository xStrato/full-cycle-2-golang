package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Id         string `gorm:"type:uuid;primaryKey"`
	Name       string
	Categories []Category `gorm:"foreignKey:Id"`
}
