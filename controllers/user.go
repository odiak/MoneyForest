package controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/store"
	"github.com/odiak/MoneyForest/util"
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

func ToUserMedia(user *store.User) *app.UserMedia {
	return &app.UserMedia{
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserMediaWithToken(user *store.User) *app.UserMediaWithToken {
	return &app.UserMediaWithToken{
		Name:  user.Name,
		Email: user.Email,
	}
}

func CreateToken(db orm.DB, userID string) (string, error) {
	token, err := util.RandomStr(80)
	if err != nil {
		return "", err
	}
	ut := &store.UserToken{
		UserID: userID,
		Token:  token,
	}
	if err := db.Insert(ut); err != nil {
		return "", err
	}
	return token, nil
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	user := store.User{}

	err := c.db.Model(&user).Where("email = ?", ctx.Payload.Email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.Unauthorized()
		}
		return c.UnexpectedError(err)
	}

	if !user.ValidPassword(ctx.Payload.Password) {
		return ctx.Unauthorized()
	}

	um := ToUserMediaWithToken(&user)
	token, err := CreateToken(c.db, user.ID)
	if err != nil {
		c.Service.LogError(err.Error())
		return goa.ErrInternal("unknown error")
	}
	um.Token = token
	return ctx.OKWithToken(um)
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
	um := ToUserMediaWithToken(u)
	token, err := CreateToken(c.db, u.ID)
	if err != nil {
		c.Service.LogError(err.Error())
		return goa.ErrInternal("unknown error")
	}
	um.Token = token
	return ctx.OKWithToken(um)
}

func (c *UserController) GetMyInfo(ctx *app.GetMyInfoUserContext) error {
	currentUser := GetCurrentUser(ctx)

	return ctx.OK(ToUserMedia(currentUser))
}
