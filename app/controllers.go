// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "MoneyForest": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Muxer
	Create(*CreateAccountContext) error
	Delete(*DeleteAccountContext) error
	List(*ListAccountContext) error
	Show(*ShowAccountContext) error
	Update(*UpdateAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service *goa.Service, ctrl AccountController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("POST", "/api/accounts", ctrl.MuxHandler("create", h, unmarshalCreateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Create", "route", "POST /api/accounts", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("DELETE", "/api/accounts/:accountID", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Delete", "route", "DELETE /api/accounts/:accountID", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("GET", "/api/accounts", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "List", "route", "GET /api/accounts", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("GET", "/api/accounts/:accountID", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Show", "route", "GET /api/accounts/:accountID", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("PUT", "/api/accounts/:accountID", ctrl.MuxHandler("update", h, unmarshalUpdateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Update", "route", "PUT /api/accounts/:accountID", "security", "APIKeyAuth")
}

// unmarshalCreateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &accountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &accountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// CategoryController is the controller interface for the Category actions.
type CategoryController interface {
	goa.Muxer
	Create(*CreateCategoryContext) error
	Delete(*DeleteCategoryContext) error
	List(*ListCategoryContext) error
	Update(*UpdateCategoryContext) error
}

// MountCategoryController "mounts" a Category resource controller on the given service.
func MountCategoryController(service *goa.Service, ctrl CategoryController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CategoryPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("POST", "/api/categories", ctrl.MuxHandler("create", h, unmarshalCreateCategoryPayload))
	service.LogInfo("mount", "ctrl", "Category", "action", "Create", "route", "POST /api/categories", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("DELETE", "/api/categories/:categoryID", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Category", "action", "Delete", "route", "DELETE /api/categories/:categoryID", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("GET", "/api/categories", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Category", "action", "List", "route", "GET /api/categories", "security", "APIKeyAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateCategoryPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("APIKeyAuth", h)
	service.Mux.Handle("PUT", "/api/categories/:categoryID", ctrl.MuxHandler("update", h, unmarshalUpdateCategoryPayload))
	service.LogInfo("mount", "ctrl", "Category", "action", "Update", "route", "PUT /api/categories/:categoryID", "security", "APIKeyAuth")
}

// unmarshalCreateCategoryPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCategoryPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &categoryPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateCategoryPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCategoryPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateCategoryPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Login(*LoginUserContext) error
	Register(*RegisterUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	service.Mux.Handle("POST", "/api/users/login", ctrl.MuxHandler("login", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "POST /api/users/login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Register(rctx)
	}
	service.Mux.Handle("POST", "/api/users", ctrl.MuxHandler("register", h, unmarshalRegisterUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Register", "route", "POST /api/users")
}

// unmarshalRegisterUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
