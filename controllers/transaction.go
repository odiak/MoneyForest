package controllers

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/constants"
	"github.com/odiak/MoneyForest/store"
	"github.com/odiak/MoneyForest/util"
	uuid "github.com/satori/go.uuid"
)

type TransactionController struct {
	*CommonController
	db orm.DB
}

// NewTransactionController creates a user controller.
func NewTransactionController(service *goa.Service, db orm.DB) *TransactionController {
	return &TransactionController{
		CommonController: NewCommonController(service, "TransactionController"),
		db:               db,
	}
}

func ToTransactionMedia(transaction *store.Transaction) *app.TransactionMedia {
	id, _ := uuid.FromString(transaction.ID)
	accountID, _ := uuid.FromString(transaction.AccountID)
	return &app.TransactionMedia{
		ID:              id,
		AccountID:       accountID,
		Amount:          int(transaction.Amount),
		TransactionType: transaction.TransactionType,
		Title:           transaction.Title,
		OriginalTitle:   transaction.OriginalTitle,
		Description:     transaction.Description,
		Category:        ToCategoryMediaWithParent(transaction.Category),
		Date:            transaction.Date.Format("2006-01-02"),
	}
}

func ToTransactionMediaList(transactions []store.Transaction) []*app.TransactionMedia {
	slice := make([]*app.TransactionMedia, len(transactions))
	for i, transaction := range transactions {
		slice[i] = ToTransactionMedia(&transaction)
	}
	return slice
}

func FromTransactionPayload(payload *app.TransactionPayload) *store.Transaction {
	date, err := time.Parse("2006-01-02", payload.Date)
	if err != nil {
		panic(err)
	}
	var categoryID *string
	if payload.CategoryID != nil {
		categoryID = util.StringPtr(payload.CategoryID.String())
	}
	return &store.Transaction{
		AccountID:       payload.AccountID.String(),
		Amount:          int32(payload.Amount),
		TransactionType: payload.TransactionType,
		Title:           payload.Title,
		OriginalTitle:   payload.OriginalTitle,
		Description:     payload.Description,
		CategoryID:      categoryID,
		Date:            date,
	}
}

func UpdateFromTransactionPayload(transaction *store.Transaction, payload *app.TransactionPayload) {
	date, err := time.Parse("2006-01-02", payload.Date)
	if err != nil {
		panic(err)
	}
	transaction.AccountID = payload.AccountID.String()
	transaction.Amount = int32(payload.Amount)
	transaction.TransactionType = payload.TransactionType
	transaction.Title = payload.Title
	transaction.OriginalTitle = payload.OriginalTitle
	transaction.Description = payload.Description
	transaction.CategoryID = util.StringPtr(payload.CategoryID.String())
	transaction.Date = date
}

func (c *TransactionController) List(ctx *app.ListTransactionContext) error {
	currentUser := GetCurrentUser(ctx)

	count := ctx.Count
	page := ctx.Page
	var transactions []store.Transaction
	q := c.db.Model(&transactions).
		Where("owner_id = ?", currentUser.ID).
		Join("INNER JOIN accounts account ON account.id = transaction.account_id").
		Where("account.owner_id = ?", currentUser.ID)
	if ctx.AccountID != nil {
		q = q.Where("account_id = ?", ctx.AccountID.String())
	}
	err := q.Order("name").
		Limit(count + 1).
		Offset((page - 1) * count).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return c.UnexpectedError(err)
	}
	hasNext := len(transactions) > count
	if hasNext {
		transactions = transactions[:len(transactions)-1]
	}
	return ctx.OK(&app.TransactionListMedia{
		Transactions: ToTransactionMediaList(transactions),
		HasNext:      hasNext,
	})
}

func (c *TransactionController) Show(ctx *app.ShowTransactionContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	account := &store.Transaction{ID: ctx.TransactionID.String()}
	err := c.db.Model(account).
		Where("id = ?id AND owner_id = ?", currentUser.ID).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return ctx.NotFound()
		}
		return c.UnexpectedError(err)
	}
	return ctx.OK(ToTransactionMedia(account))
}

func (c *TransactionController) Create(ctx *app.CreateTransactionContext) error {
	_ = ctx.Value(constants.CurrentUserKey).(*store.User)

	transaction := FromTransactionPayload(ctx.Payload)
	err := c.db.Insert(transaction)
	if err != nil {
		return c.UnexpectedError(err)
	}

	if transaction.CategoryID != nil {
		category := &store.Category{}
		err = c.db.Model(category).Column("category.*", "ParentCategory").Where("category.id = ?", transaction.CategoryID).Select()
		if err != nil {
			return c.UnexpectedError(err)
		}
		transaction.Category = category
	}

	return ctx.OK(ToTransactionMedia(transaction))
}

func (c *TransactionController) Update(ctx *app.UpdateTransactionContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	accountID := ctx.TransactionID.String()
	account := store.Transaction{}
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

	UpdateFromTransactionPayload(&account, ctx.Payload)
	err = c.db.Update(&account)
	if err != nil {
		return c.UnexpectedError(err)
	}

	return ctx.OK(ToTransactionMedia(&account))
}

func (c *TransactionController) Delete(ctx *app.DeleteTransactionContext) error {
	currentUser := ctx.Value(constants.CurrentUserKey).(*store.User)

	account := &store.Transaction{ID: ctx.TransactionID.String()}
	_, err := c.db.Model(account).
		Where("id = ?id").
		Where("owner_id = ?", currentUser.ID).
		Delete()
	if err != nil {
		return c.UnexpectedError(err)
	}

	return ctx.NoContent()
}
