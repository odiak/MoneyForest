// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "MoneyForest": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/odiak/MoneyForest/design
// --out=$(GOPATH)/src/github.com/odiak/MoneyForest
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"unicode/utf8"
)

// AccountMedia media type (default view)
//
// Identifier: application/vnd.moneyforest.account+json; view=default
type AccountMedia struct {
	AccountType string    `form:"accountType" json:"accountType" xml:"accountType"`
	Balance     int       `form:"balance" json:"balance" xml:"balance"`
	Description string    `form:"description" json:"description" xml:"description"`
	HasBalance  bool      `form:"hasBalance" json:"hasBalance" xml:"hasBalance"`
	ID          uuid.UUID `form:"id" json:"id" xml:"id"`
	Name        string    `form:"name" json:"name" xml:"name"`
}

// Validate validates the AccountMedia media type instance.
func (mt *AccountMedia) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.AccountType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "accountType"))
	}

	if !(mt.AccountType == "wallet" || mt.AccountType == "bank" || mt.AccountType == "credit-card") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.accountType`, mt.AccountType, []interface{}{"wallet", "bank", "credit-card"}))
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	return
}

// AccountListMedia media type (default view)
//
// Identifier: application/vnd.moneyforest.account-list+json; view=default
type AccountListMedia struct {
	Accounts []*AccountMedia `form:"accounts" json:"accounts" xml:"accounts"`
	HasNext  bool            `form:"hasNext" json:"hasNext" xml:"hasNext"`
}

// Validate validates the AccountListMedia media type instance.
func (mt *AccountListMedia) Validate() (err error) {
	if mt.Accounts == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "accounts"))
	}

	for _, e := range mt.Accounts {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CategoryMedia media type (default view)
//
// Identifier: application/vnd.moneyforest.category+json; view=default
type CategoryMedia struct {
	ID               uuid.UUID  `form:"id" json:"id" xml:"id"`
	Name             string     `form:"name" json:"name" xml:"name"`
	ParentCategoryID *uuid.UUID `form:"parentCategoryId,omitempty" json:"parentCategoryId,omitempty" xml:"parentCategoryId,omitempty"`
}

// Validate validates the CategoryMedia media type instance.
func (mt *CategoryMedia) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// CategoryMedia media type (full view)
//
// Identifier: application/vnd.moneyforest.category+json; view=full
type CategoryMediaFull struct {
	ID             uuid.UUID          `form:"id" json:"id" xml:"id"`
	Name           string             `form:"name" json:"name" xml:"name"`
	ParentCategory *CategoryMediaFull `form:"parentCategory,omitempty" json:"parentCategory,omitempty" xml:"parentCategory,omitempty"`
}

// Validate validates the CategoryMediaFull media type instance.
func (mt *CategoryMediaFull) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.ParentCategory != nil {
		if err2 := mt.ParentCategory.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// TransactionMedia media type (default view)
//
// Identifier: application/vnd.moneyforest.transaction+json; view=default
type TransactionMedia struct {
	AccountID       uuid.UUID      `form:"accountId" json:"accountId" xml:"accountId"`
	Amount          int            `form:"amount" json:"amount" xml:"amount"`
	Category        *CategoryMedia `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	Date            string         `form:"date" json:"date" xml:"date"`
	Description     string         `form:"description" json:"description" xml:"description"`
	ID              uuid.UUID      `form:"id" json:"id" xml:"id"`
	OriginalTitle   string         `form:"originalTitle" json:"originalTitle" xml:"originalTitle"`
	Title           string         `form:"title" json:"title" xml:"title"`
	TransactionType string         `form:"transactionType" json:"transactionType" xml:"transactionType"`
}

// Validate validates the TransactionMedia media type instance.
func (mt *TransactionMedia) Validate() (err error) {

	if mt.TransactionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "transactionType"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.OriginalTitle == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "originalTitle"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.Date == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "date"))
	}
	if mt.Category != nil {
		if err2 := mt.Category.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`^\d{1,4}-\d{2}-\d{2}$`, mt.Date); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.date`, mt.Date, `^\d{1,4}-\d{2}-\d{2}$`))
	}
	if !(mt.TransactionType == "expense" || mt.TransactionType == "income" || mt.TransactionType == "transfer" || mt.TransactionType == "balance-adjustment") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.transactionType`, mt.TransactionType, []interface{}{"expense", "income", "transfer", "balance-adjustment"}))
	}
	return
}

// user information (default view)
//
// Identifier: application/vnd.moneyforest.user+json; view=default
type UserMedia struct {
	Email string `form:"email" json:"email" xml:"email"`
	Name  string `form:"name" json:"name" xml:"name"`
}

// Validate validates the UserMedia media type instance.
func (mt *UserMedia) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	return
}
