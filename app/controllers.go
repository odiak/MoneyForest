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
	service.Mux.Handle("PUT", "/api/users", ctrl.MuxHandler("login", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "PUT /api/users")

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
			rctx.Payload = rawPayload.(*RegisterUserPayload)
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
	payload := &registerUserPayload{}
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
