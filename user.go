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
	db pg.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db *pg.DB) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	return nil
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	return nil
}
