package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserMedia = MediaType("application/vnd.moneyforest.user+json", func() {
	TypeName("UserMedia")
	Description("user information")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("name")
		Attribute("email")

		Required("name", "email")
	})

	View("default", func() {
		Attribute("name")
		Attribute("email")
	})
})

var AccountMedia = MediaType("application/vnd.moneyforest.account+json", func() {
	TypeName("AccountMedia")
	Reference(AccountPayload)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("name")
		Attribute("description")
		Attribute("accountType")
		Attribute("hasBalance")
		Attribute("balance")

		Required("id", "name", "description", "accountType", "hasBalance", "balance")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("accountType")
		Attribute("hasBalance")
		Attribute("balance")
	})
})

var AccountListMedia = MediaType("application/vnd.moneyforest.account-list+json", func() {
	TypeName("AccountListMedia")

	Attributes(func() {
		Attribute("accounts", ArrayOf(AccountMedia))
		Attribute("hasNext", Boolean)

		Required("accounts", "hasNext")
	})

	View("default", func() {
		Attribute("accounts")
		Attribute("hasNext")
	})
})

var TransactionMedia = MediaType("application/vnd.moneyforest.transaction+json", func() {
	TypeName("TransactionMedia")
	Reference(TransactionPayload)
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("accountId")
		Attribute("amount")
		Attribute("transactionType")
		Attribute("title")
		Attribute("originalTitle")
		Attribute("description")
		Attribute("category", CategoryMedia, func() {
			View("withParent")
		})
		Attribute("date")

		Required("id", "accountId", "amount", "transactionType", "title", "date")
	})

	View("default", func() {
		Attribute("id")
		Attribute("accountId")
		Attribute("amount")
		Attribute("transactionType")
		Attribute("title")
		Attribute("originalTitle")
		Attribute("description")
		Attribute("category")
		Attribute("date")
	})
})

var TransactionListMedia = MediaType("application/vnd.moneyforest.transaction-list+json", func() {
	TypeName("TransactionListMedia")

	Attributes(func() {
		Attribute("transactions", CollectionOf(TransactionMedia))
		Attribute("hasNext", Boolean)

		Required("transactions", "hasNext")
	})

	View("default", func() {
		Attribute("transactions")
		Attribute("hasNext")
	})
})

var CategoryMedia = MediaType("application/vnd.moneyforest.category+json", func() {
	TypeName("CategoryMedia")
	Reference(CategoryPayload)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("name")
		Attribute("parentCategoryId", UUID)
		Attribute("parentCategory", "application/vnd.moneyforest.category")
		Attribute("childCategories", CollectionOf("application/vnd.moneyforest.category"))

		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("parentCategoryId")
	})

	View("withParent", func() {
		Attribute("id")
		Attribute("name")
		Attribute("parentCategory", func() {
			View("withParent")
		})
	})

	View("withChildren", func() {
		Attribute("id")
		Attribute("name")
		Attribute("parentCategoryId")
		Attribute("childCategories", func() {
			View("withChildren")
		})
	})
})

var CategoryListMedia = MediaType("application/vnd.moneyforest.category-list+json", func() {
	TypeName("CategoryListMedia")

	Attributes(func() {
		Attribute("categories", CollectionOf(CategoryMedia))
		Attribute("hasNext", Boolean)

		Required("categories", "hasNext")
	})

	View("default", func() {
		Attribute("categories", func() {
			View("withChildren")
		})
		Attribute("hasNext")
	})
})
