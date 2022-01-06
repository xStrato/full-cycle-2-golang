package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/webapi/controllers"
)

type CategoryRouter struct {
	ctr *controllers.CategoryController
}

func NewCategoryRouter(c *controllers.CategoryController) *CategoryRouter {
	return &CategoryRouter{c}
}

func (cr *CategoryRouter) Configure(gin *gin.Engine) error {
	main := gin.Group("api/v1")
	{
		category := main.Group("category")
		{
			category.POST("/", cr.ctr.Create)
		}
	}
	return nil
}
