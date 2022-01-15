package models

import (
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/constants"
	"gorm.io/gorm"
)

type CastMember struct {
	gorm.Model
	Id             string `gorm:"type:uuid;primaryKey"`
	Name           string
	CastMemberType constants.CastMemberType
}
