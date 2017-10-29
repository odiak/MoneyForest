package controllers

import (
	"github.com/goadesign/goa"
)

var (
	duplicatedEmailErr = goa.NewErrorClass("duplicated_email", 1000)
)

func unexpectedError(service *goa.Service, err error) error {
	service.LogError("Unexpected error", "err", err)
	return goa.ErrInternal("unexpected error")
}
