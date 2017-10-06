package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("MoneyForest", func() {
	Title("money management system")

	Host("localhost:8000")
	Scheme("http")
	BasePath("/api")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("user", func() {
	DefaultMedia(UserMedia)
	BasePath("/users")

	Action("register", func() {
		Routing(
			POST(""),
		)
		Payload(UserPayload, func() {
			Required("name", "email", "password")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("login", func() {
		Routing(
			PUT(""),
		)
		Params(func() {
			Param("email", String)
			Param("password", String)
		})
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

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
})

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("user information")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("name")
		Attribute("email")
	})

	View("default", func() {
		Attribute("name")
		Attribute("email")
	})
})
