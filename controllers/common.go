package controllers

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/constants"
	"github.com/odiak/MoneyForest/store"
)

type CommonController struct {
	*goa.Controller
}

func NewCommonController(service *goa.Service, name string) *CommonController {
	return &CommonController{service.NewController(name)}
}

func (c *CommonController) UnexpectedError(err error) error {
	c.Service.LogError("Unexpected error", "err", err)
	return goa.ErrInternal("unexpected error")
}

func (c *CommonController) CurrentUser() *store.User {
	return c.Context.Value(constants.CurrentUserKey).(*store.User)
}

func GetCurrentUser(ctx context.Context) *store.User {
	return ctx.Value(constants.CurrentUserKey).(*store.User)
}
