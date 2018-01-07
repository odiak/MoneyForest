package controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/store"
	uuid "github.com/satori/go.uuid"
)

type CategoryController struct {
	*CommonController
	db orm.DB
}

func NewCategoryController(service *goa.Service, db orm.DB) *CategoryController {
	return &CategoryController{
		CommonController: NewCommonController(service, "CategoryController"),
		db:               db,
	}
}

func FromCategoryPayload(payload *app.CategoryPayload) *store.Category {
	var parentID *string
	if payload.ParentCategoryID != nil {
		id := payload.ParentCategoryID.String()
		parentID = &id
	}
	return &store.Category{
		Name:             payload.Name,
		ParentCategoryID: parentID,
	}
}

func ToCategoryMedia(category *store.Category) *app.CategoryMedia {
	if category == nil {
		return nil
	}

	ID, _ := uuid.FromString(category.ID)
	var parentID *uuid.UUID
	if category.ParentCategoryID != nil {
		pid, _ := uuid.FromString(*category.ParentCategoryID)
		parentID = &pid
	}
	return &app.CategoryMedia{
		ID:               ID,
		Name:             category.Name,
		ParentCategoryID: parentID,
	}
}

func ToCategoryMediaWithChildren(category *store.Category) *app.CategoryMediaWithChildren {
	if category == nil {
		return nil
	}

	ID, _ := uuid.FromString(category.ID)
	var parentID *uuid.UUID
	if category.ParentCategoryID != nil {
		pid, _ := uuid.FromString(*category.ParentCategoryID)
		parentID = &pid
	}
	var children []*app.CategoryMediaWithChildren
	if category.ChildCategories != nil && len(category.ChildCategories) > 0 {
		children = ToCategoryMediaWithChildrenList(category.ChildCategories)
	}
	return &app.CategoryMediaWithChildren{
		ID:               ID,
		Name:             category.Name,
		ParentCategoryID: parentID,
		ChildCategories:  children,
	}
}

func ToCategoryMediaWithParent(category *store.Category) *app.CategoryMediaWithParent {
	if category == nil {
		return nil
	}

	var parent *app.CategoryMediaWithParent
	if category.ParentCategory != nil {
		parent = ToCategoryMediaWithParent(category.ParentCategory)
	}
	return &app.CategoryMediaWithParent{
		ID:             uuid.FromStringOrNil(category.ID),
		Name:           category.Name,
		ParentCategory: parent,
	}
}

func ToCategoryMediaWithChildrenList(categories []store.Category) []*app.CategoryMediaWithChildren {
	slice := make([]*app.CategoryMediaWithChildren, len(categories))
	for i, category := range categories {
		slice[i] = ToCategoryMediaWithChildren(&category)
	}
	return slice
}

func (c *CategoryController) List(ctx *app.ListCategoryContext) error {
	currentUser := GetCurrentUser(ctx)

	count := ctx.Count
	page := ctx.Page

	var categories []store.Category
	err := c.db.Model(&categories).Column("category.*", "ChildCategories").
		Where("owner_id = ? AND parent_category_id IS NULL", currentUser.ID).
		Order("name").
		Limit(count + 1).
		Offset((page - 1) * count).
		Select()
	if err != nil {
		return c.UnexpectedError(err)
	}
	hasNext := len(categories) > count
	if hasNext {
		categories = categories[:len(categories)-1]
	}
	return ctx.OK(&app.CategoryListMedia{
		Categories: ToCategoryMediaWithChildrenList(categories),
		HasNext:    hasNext,
	})
}

func (c *CategoryController) Create(ctx *app.CreateCategoryContext) error {
	currentUser := GetCurrentUser(ctx)

	category := FromCategoryPayload(ctx.Payload)
	if category.ParentCategoryID != nil {
		parent := store.Category{ID: *category.ParentCategoryID}
		err := c.db.Model(&parent).Where("id = ?id AND owner_id = ?", currentUser.ID).
			Limit(1).Select()
		if err != nil {
			return c.UnexpectedError(err)
		}
		if parent.ParentCategoryID != nil {
			return goa.ErrBadRequest("parent must not have parent")
		}
	}
	category.OwnerID = currentUser.ID
	err := c.db.Insert(category)
	if err != nil {
		return c.UnexpectedError(err)
	}
	return ctx.OK(ToCategoryMedia(category))
}

func (c *CategoryController) Update(ctx *app.UpdateCategoryContext) error {
	currentUser := GetCurrentUser(ctx)

	categoryID := ctx.CategoryID.String()
	category := store.Category{}
	err := c.db.Model(&category).
		Where("id = ? AND owner_id = ?", categoryID, currentUser.ID).
		Limit(1).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.NotFound()
		}
		return c.UnexpectedError(err)
	}

	category.Name = ctx.Payload.Name
	err = c.db.Update(&category)
	if err != nil {
		c.UnexpectedError(err)
	}

	return ctx.OK(ToCategoryMedia(&category))
}

func (c *CategoryController) Delete(ctx *app.DeleteCategoryContext) error {
	currentUser := GetCurrentUser(ctx)

	id := ctx.CategoryID.String()

	_, err := c.db.Model(&store.Category{}).
		Where("parent_category_id = ?", id).
		Delete()
	if err != nil {
		return c.UnexpectedError(err)
	}
	_, err = c.db.Model(&store.Category{}).
		Where("id = ? AND owner_id = ?", id, currentUser.ID).
		Delete()
	if err != nil {
		return c.UnexpectedError(err)
	}

	return ctx.NoContent()
}
