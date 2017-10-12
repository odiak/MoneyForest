package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserPayload = Type("UserPayload", func() {
	Attribute("name", func() {
		MinLength(2)
		Example("James Brown")
	})
	Attribute("email", func() {
		Format("email")
	})
	Attribute("password", func() {
		MinLength(8)
	})

	Required("name", "email", "password")
})

var AccountPayload = Type("AccountPayload", func() {
	Attribute("name", func() {
		MinLength(1)
	})
	Attribute("description", func() {
		Default("")
	})
	Attribute("accountType", func() {
		Enum("wallet", "bank", "credit-card")
	})
	Attribute("hasBalance", Boolean, func() {
		Default(false)
	})
	Attribute("balance", Integer, func() {
		Default(0)
	})

	Required("name", "description", "accountType", "hasBalance", "balance")
})

var TransactionPayload = Type("TransactionPayload", func() {
	Attribute("accountId", UUID)
	Attribute("amount", Integer)
	Attribute("transactionType", func() {
		Enum("expense", "income", "transfer", "balance-adjustment")
	})
	Attribute("title")
	Attribute("originalTitle", func() {
		Default("")
	})
	Attribute("description", func() {
		Default("")
	})
	Attribute("categoryId", UUID)
	Attribute("date", func() {
		Pattern(`^\d{1,4}-\d{2}-\d{2}$`)
	})

	Required("accountId", "amount", "transactionType", "title", "date")
})

var CategoryPayload = Type("Category", func() {
	Attribute("name")
	Attribute("parentCategoryId", UUID)

	Required("name")
})
