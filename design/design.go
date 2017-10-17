package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var APIKeyAuth = APIKeySecurity("APIKeyAuth", func() {
	Header("X-MoneyForest-Auth-Token")
})

var _ = API("MoneyForest", func() {
	Title("money management system")

	Host("localhost:8000")
	Scheme("http")
	BasePath("/api")
	Consumes("application/json")
	Produces("application/json")

	Security(APIKeyAuth)
})

var _ = Resource("user", func() {
	DefaultMedia(UserMedia)
	BasePath("/users")
	NoSecurity()

	Action("register", func() {
		Routing(
			POST(""),
		)
		Payload(UserPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("login", func() {
		Routing(
			POST("/login"),
		)
		Params(func() {
			Param("email", String)
			Param("password", String)

			Required("email", "password")
		})
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("account", func() {
	DefaultMedia(AccountMedia)
	BasePath("/accounts")

	Action("create", func() {
		Routing(
			POST(""),
		)
		Payload(AccountPayload)
		Response(OK)
	})

	Action("show", func() {
		Routing(
			GET(":accountID"),
		)
		Params(func() {
			Param("accountID", UUID)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("list", func() {
		Routing(
			GET(""),
		)
		Params(func() {
			Param("count", Integer, func() {
				Minimum(1)
				Maximum(60)
				Default(30)
			})
			Param("page", Integer, func() {
				Minimum(1)
				Default(1)
			})
		})
		Response(OK, AccountListMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT(":accountID"),
		)
		Params(func() {
			Param("accountID", UUID)
		})
		Payload(AccountPayload)
		Response(OK, AccountMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE(":accountID"),
		)
		Params(func() {
			Param("accountID", UUID)
		})
		Response(NoContent)
		Response(NotFound)
	})
})

var _ = Resource("category", func() {
	DefaultMedia(CategoryMedia)
	BasePath("/categories")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Params(func() {
			Param("count", Integer, func() {
				Minimum(1)
				Maximum(60)
				Default(40)
			})
			Param("page", Integer, func() {
				Minimum(1)
				Default(1)
			})
		})
		Response(OK, CategoryListMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Payload(CategoryPayload)
		Response("OK", CategoryMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT(":categoryID"),
		)
		Params(func() {
			Param("categoryID", UUID)
		})
		Payload(func() {
			Member("name", String, func() {
				MinLength(1)
			})
			Required("name")
		})
		Response("OK", CategoryMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE(":categoryID"),
		)
		Params(func() {
			Param("categoryID", UUID)
		})
		Response(NoContent)
		Response(NotFound)
	})
})
