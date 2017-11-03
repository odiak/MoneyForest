// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "MoneyForest": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/odiak/MoneyForest/design
// --out=$(GOPATH)/src/github.com/odiak/MoneyForest
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"unicode/utf8"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateAccountContext) OK(r *AccountMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID uuid.UUID
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := uuid.FromString(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "uuid"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListAccountContext provides the account list action context.
type ListAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Count int
	Page  int
}

// NewListAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller list action.
func NewListAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCount := req.Params["count"]
	if len(paramCount) == 0 {
		rctx.Count = 30
	} else {
		rawCount := paramCount[0]
		if count, err2 := strconv.Atoi(rawCount); err2 == nil {
			rctx.Count = count
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("count", rawCount, "integer"))
		}
		if rctx.Count < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 1, true))
		}
		if rctx.Count > 60 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 60, false))
		}
	}
	paramPage := req.Params["page"]
	if len(paramPage) == 0 {
		rctx.Page = 1
	} else {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			rctx.Page = page
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, rctx.Page, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OK(r *AccountListMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.account-list+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID uuid.UUID
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := uuid.FromString(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(r *AccountMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID uuid.UUID
	Payload   *AccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := uuid.FromString(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateAccountContext) OK(r *AccountMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateCategoryContext provides the category create action context.
type CreateCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CategoryPayload
}

// NewCreateCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller create action.
func NewCreateCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCategoryContext) OK(r *CategoryMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKWithChildren sends a HTTP response with status code 200.
func (ctx *CreateCategoryContext) OKWithChildren(r *CategoryMediaWithChildren) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKWithParent sends a HTTP response with status code 200.
func (ctx *CreateCategoryContext) OKWithParent(r *CategoryMediaWithParent) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCategoryContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// DeleteCategoryContext provides the category delete action context.
type DeleteCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CategoryID uuid.UUID
}

// NewDeleteCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller delete action.
func NewDeleteCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCategoryID := req.Params["categoryID"]
	if len(paramCategoryID) > 0 {
		rawCategoryID := paramCategoryID[0]
		if categoryID, err2 := uuid.FromString(rawCategoryID); err2 == nil {
			rctx.CategoryID = categoryID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("categoryID", rawCategoryID, "uuid"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteCategoryContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteCategoryContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListCategoryContext provides the category list action context.
type ListCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Count int
	Page  int
}

// NewListCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller list action.
func NewListCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCount := req.Params["count"]
	if len(paramCount) == 0 {
		rctx.Count = 40
	} else {
		rawCount := paramCount[0]
		if count, err2 := strconv.Atoi(rawCount); err2 == nil {
			rctx.Count = count
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("count", rawCount, "integer"))
		}
		if rctx.Count < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 1, true))
		}
		if rctx.Count > 60 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 60, false))
		}
	}
	paramPage := req.Params["page"]
	if len(paramPage) == 0 {
		rctx.Page = 1
	} else {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			rctx.Page = page
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, rctx.Page, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCategoryContext) OK(r *CategoryListMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category-list+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// UpdateCategoryContext provides the category update action context.
type UpdateCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CategoryID uuid.UUID
	Payload    *UpdateCategoryPayload
}

// NewUpdateCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller update action.
func NewUpdateCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCategoryID := req.Params["categoryID"]
	if len(paramCategoryID) > 0 {
		rawCategoryID := paramCategoryID[0]
		if categoryID, err2 := uuid.FromString(rawCategoryID); err2 == nil {
			rctx.CategoryID = categoryID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("categoryID", rawCategoryID, "uuid"))
		}
	}
	return &rctx, err
}

// updateCategoryPayload is the category update action payload.
type updateCategoryPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateCategoryPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 1, true))
		}
	}
	return
}

// Publicize creates UpdateCategoryPayload from updateCategoryPayload
func (payload *updateCategoryPayload) Publicize() *UpdateCategoryPayload {
	var pub UpdateCategoryPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateCategoryPayload is the category update action payload.
type UpdateCategoryPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateCategoryPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if utf8.RuneCountInString(payload.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, utf8.RuneCountInString(payload.Name), 1, true))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateCategoryContext) OK(r *CategoryMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKWithChildren sends a HTTP response with status code 200.
func (ctx *UpdateCategoryContext) OKWithChildren(r *CategoryMediaWithChildren) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKWithParent sends a HTTP response with status code 200.
func (ctx *UpdateCategoryContext) OKWithParent(r *CategoryMediaWithParent) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateCategoryContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateCategoryContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateTransactionContext provides the transaction create action context.
type CreateTransactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *TransactionPayload
}

// NewCreateTransactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the transaction controller create action.
func NewCreateTransactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateTransactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateTransactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateTransactionContext) OK(r *TransactionMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.transaction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateTransactionContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// DeleteTransactionContext provides the transaction delete action context.
type DeleteTransactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TransactionID uuid.UUID
}

// NewDeleteTransactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the transaction controller delete action.
func NewDeleteTransactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteTransactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteTransactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTransactionID := req.Params["transactionID"]
	if len(paramTransactionID) > 0 {
		rawTransactionID := paramTransactionID[0]
		if transactionID, err2 := uuid.FromString(rawTransactionID); err2 == nil {
			rctx.TransactionID = transactionID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("transactionID", rawTransactionID, "uuid"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteTransactionContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteTransactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListTransactionContext provides the transaction list action context.
type ListTransactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID *uuid.UUID
	Count     int
	Page      int
}

// NewListTransactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the transaction controller list action.
func NewListTransactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListTransactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListTransactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := uuid.FromString(rawAccountID); err2 == nil {
			tmp11 := &accountID
			rctx.AccountID = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "uuid"))
		}
	}
	paramCount := req.Params["count"]
	if len(paramCount) == 0 {
		rctx.Count = 40
	} else {
		rawCount := paramCount[0]
		if count, err2 := strconv.Atoi(rawCount); err2 == nil {
			rctx.Count = count
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("count", rawCount, "integer"))
		}
		if rctx.Count < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 1, true))
		}
		if rctx.Count > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`count`, rctx.Count, 100, false))
		}
	}
	paramPage := req.Params["page"]
	if len(paramPage) == 0 {
		rctx.Page = 1
	} else {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			rctx.Page = page
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, rctx.Page, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTransactionContext) OK(r *TransactionListMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.transaction-list+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowTransactionContext provides the transaction show action context.
type ShowTransactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TransactionID uuid.UUID
}

// NewShowTransactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the transaction controller show action.
func NewShowTransactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowTransactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowTransactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTransactionID := req.Params["transactionID"]
	if len(paramTransactionID) > 0 {
		rawTransactionID := paramTransactionID[0]
		if transactionID, err2 := uuid.FromString(rawTransactionID); err2 == nil {
			rctx.TransactionID = transactionID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("transactionID", rawTransactionID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTransactionContext) OK(r *TransactionMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.transaction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTransactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateTransactionContext provides the transaction update action context.
type UpdateTransactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TransactionID uuid.UUID
	Payload       *TransactionPayload
}

// NewUpdateTransactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the transaction controller update action.
func NewUpdateTransactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateTransactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateTransactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTransactionID := req.Params["transactionID"]
	if len(paramTransactionID) > 0 {
		rawTransactionID := paramTransactionID[0]
		if transactionID, err2 := uuid.FromString(rawTransactionID); err2 == nil {
			rctx.TransactionID = transactionID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("transactionID", rawTransactionID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateTransactionContext) OK(r *TransactionMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.transaction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateTransactionContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateTransactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// LoginUserContext provides the user login action context.
type LoginUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *LoginUserPayload
}

// NewLoginUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller login action.
func NewLoginUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// loginUserPayload is the user login action payload.
type loginUserPayload struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *loginUserPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates LoginUserPayload from loginUserPayload
func (payload *loginUserPayload) Publicize() *LoginUserPayload {
	var pub LoginUserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// LoginUserPayload is the user login action payload.
type LoginUserPayload struct {
	Email    string `form:"email" json:"email" xml:"email"`
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *LoginUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginUserContext) OK(r *UserMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *LoginUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// RegisterUserContext provides the user register action context.
type RegisterUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UserPayload
}

// NewRegisterUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller register action.
func NewRegisterUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*RegisterUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RegisterUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RegisterUserContext) OK(r *UserMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.moneyforest.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *RegisterUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RegisterUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
