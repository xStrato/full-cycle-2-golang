package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/infrastructure/data/contexts"
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

	fmt.Printf("Command Type: %v\n", createCmd.GetCommandType())
	fmt.Printf("Creation Time: %v\n", createCmd.GetCreationTime().Format(time.RFC3339))
	fmt.Printf("Category Name: %v\n", createCmd.Name)

	if err := createCmd.IsValid(); err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v is a valid command.\n\n", createCmd.GetCommandType())

	categoryRepository := new(interface{})
	categoryHandler := handlers.NewCategoryCommandHandler(categoryRepository)

	result := categoryHandler.Handle(createCmd)
	fmt.Printf("Command Result - Success: %v \n", result.HasSuccess())
	fmt.Printf("Command Result - Message: %v \n", result.GetMessage())
	fmt.Printf("Command Result - Data: %v \n", result.GetData())
}
