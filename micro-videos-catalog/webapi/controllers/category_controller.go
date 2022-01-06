package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type CategoryController struct {
	categoryHandler *handlers.CategoryCommandHandler
}

func NewCategoryController(cmdHandler *handlers.CategoryCommandHandler) *CategoryController {
	return &CategoryController{
		categoryHandler: cmdHandler,
	}
}

func (ctr *CategoryController) Create(ctx *gin.Context) {
	cmd := commands.NewCreateCategoryCommand()
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Cannot bind JSON from body", err.Error()))
		return
	}
	result := ctr.categoryHandler.Handle(cmd)
	if !result.HasSuccess() {
		ctx.JSON(http.StatusBadRequest, result)
	}
	ctx.JSONP(http.StatusCreated, result)
}
