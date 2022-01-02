package contexts

import (
	"log"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/models"
	"gorm.io/gorm"
)

type VideoCatalogContext struct {
	db *gorm.DB
}

func NewVideoCatalogContext(db *gorm.DB) *VideoCatalogContext {
	return &VideoCatalogContext{db}
}

func (v *VideoCatalogContext) RunMigrations() {
	if err := v.db.AutoMigrate(&models.Category{}); err != nil {
		log.Fatalln("Cannot run migrations for Category entity: ", err.Error())
	}
	if err := v.db.AutoMigrate(&models.Genre{}); err != nil {
		log.Fatalln("Cannot run migrations for Genre entity: ", err.Error())
	}
	if err := v.db.AutoMigrate(&models.CastMember{}); err != nil {
		log.Fatalln("Cannot run migrations for CastMember entity: ", err.Error())
	}
	if err := v.db.AutoMigrate(&models.Video{}); err != nil {
		log.Fatalln("Cannot run migrations for Video entity: ", err.Error())
	}
	if err := v.db.AutoMigrate(&models.VideoFile{}); err != nil {
		log.Fatalln("Cannot run migrations for VideoFile entity: ", err.Error())
	}
}
