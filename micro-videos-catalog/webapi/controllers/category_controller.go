package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/common"
)

type CategoryController struct {
	handler *handlers.CategoryHandler
}

func NewCategoryController(handler *handlers.CategoryHandler) *CategoryController {
	return &CategoryController{
		handler: handler,
	}
}

func (ctr *CategoryController) Create(ctx *gin.Context) {
	cmd := commands.NewCreateCategoryCommand()
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Cannot bind JSON from body", err.Error()))
		return
	}
	result := ctr.handler.Handle(cmd)
	if !result.HasSuccess() {
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	ctx.JSONP(http.StatusCreated, result)
}
