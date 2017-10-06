package main

import (
	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/store"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *pg.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db *pg.DB) *UserController {
	return &UserController{Controller: service.NewController("UserController"), db: db}
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	return nil
}

func ToUserMedia(user *store.User) *app.User {
	return &app.User{
		Name:  &user.Name,
		Email: &user.Email,
	}
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	p := ctx.Payload
	u := &store.User{
		Name:  p.Name,
		Email: p.Email,
	}
	u.SetPassword(p.Password)
	err := c.db.Insert(u)
	if err != nil {
		if pgerr, ok := err.(pg.Error); ok {
			if pgerr.Field('C') == "23505" {
				return ctx.BadRequest(duplicatedEmailErr(""))
			}
		}
		c.Service.LogError(err.Error())
		return goa.ErrInternal("unknown error")
	}
	return ctx.OK(ToUserMedia(u))
}
