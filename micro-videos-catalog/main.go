package main

import (
	"log"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/contexts"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/repositories"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/webapi"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/webapi/controllers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/webapi/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	videoCatalogDbContext := contexts.NewVideoCatalogContext(db)
	videoCatalogDbContext.RunMigrations()

	categoryRepository := repositories.NewCategoryRepository()
	categoryHandler := handlers.NewCategoryCommandHandler(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryHandler)
	categoryRouter := routes.NewCategoryRouter(categoryController)

	server := webapi.NewServer("3000")
	server.Router(categoryRouter.Configure)

	server.Run()
}
