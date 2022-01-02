package main

import (
	"fmt"
	"log"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/contexts"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/repositories"
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

	createCmd := commands.NewCreateCategoryCommand("Movie")

	categoryRepository := repositories.NewCategoryRepository()
	categoryHandler := handlers.NewCategoryCommandHandler(categoryRepository)

	result, err := categoryHandler.Handle(createCmd)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("Command Result - Success: %v \n", result.HasSuccess())
	fmt.Printf("Command Result - Message: %v \n", result.GetMessage())
	fmt.Printf("Command Result - Data: %v \n", result.GetData())
}
