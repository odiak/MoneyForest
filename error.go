package main

import (
	"github.com/goadesign/goa"
)

var (
	duplicatedEmailErr = goa.NewErrorClass("duplicated_email", 1000)
)
