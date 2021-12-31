package main

import (
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func main() {
	c := entities.NewGenre("Horror")
	c.AddCategory(entities.NewCategory("Movies"))
	c.AddCategory(nil)
	a := c.GetCategories()

	a = append(a, *entities.NewCategory("Series"))
	fmt.Printf("%v\n", c.GetCategories())
	a[0] = *entities.NewCategory("Films")

	fmt.Printf("%v\n", c.GetCategories())
	fmt.Printf("%v\n", a)
}
