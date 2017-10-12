package controllers

import (
	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
)

type CategoryController struct {
	*CommonController
	db *pg.DB
}

func NewCategoryController(service *goa.Service, db *pg.DB) *CategoryController {
	return &CategoryController{
		CommonController: NewCommonController(service, "CategoryController"),
		db:               db,
	}
}

func (c *CategoryController) List(ctx *app.ListCategoryContext) error {
	return nil
}

func (c *CategoryController) Create(ctx *app.CreateCategoryContext) error {
	return nil
}

func (c *CategoryController) Update(ctx *app.UpdateCategoryContext) error {
	return nil
}

func (c *CategoryController) Delete(ctx *app.DeleteCategoryContext) error {
	return nil
}
