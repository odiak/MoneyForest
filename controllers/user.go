package controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/store"
)

// UserController implements the user resource.
type UserController struct {
	*CommonController
	db orm.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db orm.DB) *UserController {
	return &UserController{
		CommonController: NewCommonController(service, "UserController"),
		db:               db,
	}
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	user := store.User{}

	err := c.db.Model(&user).Where("email = ?", ctx.Email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.Unauthorized()
		}
		return c.UnexpectedError(err)
	}

	if !user.ValidPassword(ctx.Password) {
		return ctx.Unauthorized()
	}

	return ctx.OK(ToUserMedia(&user))
}

func ToUserMedia(user *store.User) *app.UserMedia {
	return &app.UserMedia{
		Name:  user.Name,
		Email: user.Email,
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
