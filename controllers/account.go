package controllers

import (
	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/constants"
	"github.com/odiak/MoneyForest/store"
	"github.com/satori/go.uuid"
)

// AccountController implements the user resource.
type AccountController struct {
	*CommonController
	db *pg.DB
}

// NewAccountController creates a user controller.
func NewAccountController(service *goa.Service, db *pg.DB) *AccountController {
	return &AccountController{
		CommonController: NewCommonController(service, "AccountController"),
		db:               db,
	}
}

func ToAccountMedia(account *store.Account) *app.AccountMedia {
	id, _ := uuid.FromString(account.ID)
	return &app.AccountMedia{
		ID:          id,
		Name:        account.Name,
		Description: account.Description,
		AccountType: account.AccountType,
		Balance:     int(account.Balance),
		HasBalance:  account.HasBalance,
	}
}

func ToAccountMediaList(accounts []store.Account) []*app.AccountMedia {
	slice := make([]*app.AccountMedia, len(accounts))
	for i, account := range accounts {
		slice[i] = ToAccountMedia(&account)
	}
	return slice
}

func FromAccountPayload(payload *app.AccountPayload) *store.Account {
	return &store.Account{
		Name:        payload.Name,
		Description: payload.Description,
		AccountType: payload.AccountType,
		HasBalance:  payload.HasBalance,
		Balance:     int32(payload.Balance),
	}
}

func UpdateFromAccountPayload(account *store.Account, payload *app.AccountPayload) {
	account.Name = payload.Name
	account.Description = payload.Description
	account.AccountType = payload.AccountType
	account.HasBalance = payload.HasBalance
	account.Balance = int32(payload.Balance)
}

func (c *AccountController) List(ctx *app.ListAccountContext) error {
	currentUser := GetCurrentUser(ctx)

	count := ctx.Count
	page := ctx.Page
	var accounts []store.Account
	err := c.db.Model(&accounts).
		Where("owner_id = ?", currentUser.ID).
		Order("name").
		Limit(count + 1).
		Offset((page - 1) * count).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return c.UnexpectedError(err)
	}
	hasNext := len(accounts) > count
	if hasNext {
		accounts = accounts[:len(accounts)-1]
	}
	return ctx.OK(&app.AccountListMedia{
		Accounts: ToAccountMediaList(accounts),
		HasNext:  hasNext,
	})
}

func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	account := &store.Account{ID: ctx.AccountID.String()}
	err := c.db.Model(account).
		Where("id = ?id AND owner_id = ?", currentUser.ID).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.NotFound()
		}
		return c.UnexpectedError(err)
	}
	return ctx.OK(ToAccountMedia(account))
}

func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	account := FromAccountPayload(ctx.Payload)
	account.OwnerID = currentUser.ID
	err := c.db.Insert(account)
	if err != nil {
		return c.UnexpectedError(err)
	}
	return ctx.OK(ToAccountMedia(account))
}

func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	accountID := ctx.AccountID.String()
	account := store.Account{}
	err := c.db.Model(&account).
		Where("id = ?", accountID).
		Where("owner_id = ?", currentUser.ID).
		Limit(1).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.NotFound()
		}
		return c.UnexpectedError(err)
	}

	UpdateFromAccountPayload(&account, ctx.Payload)
	err = c.db.Update(&account)
	if err != nil {
		return c.UnexpectedError(err)
	}

	return ctx.OK(ToAccountMedia(&account))
}

func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	account := &store.Account{ID: ctx.AccountID.String()}
	_, err := c.db.Model(account).
		Where("id = ?id").
		Where("owner_id = ?", currentUser.ID).
		Delete()
	if err != nil {
		return c.UnexpectedError(err)
	}

	return ctx.NoContent()
}
