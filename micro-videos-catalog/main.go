package main

import (
	"log"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/contexts"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	videoContext := contexts.NewVideoCatalogContext(db)
	videoContext.RunMigrations()

	println("running")
}
