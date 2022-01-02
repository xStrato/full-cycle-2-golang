package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id   string `gorm:"type:uuid;primaryKey"`
	Name string
}
