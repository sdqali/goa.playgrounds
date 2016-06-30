package main

import (
	"github.com/goadesign/goa"
	"playgrounds/app"
	"fmt"
)

// FieldController implements the field resource.
type FieldController struct {
	*goa.Controller
}

// NewFieldController creates a field controller.
func NewFieldController(service *goa.Service) *FieldController {
	return &FieldController{Controller: service.NewController("FieldController")}
}

// Show runs the show action.
func (c *FieldController) Show(ctx *app.ShowFieldContext) error {
	// TBD: implement

	if ctx.FieldID == "2c6cc8fb37384a12be84757786c374e5" {
		return ctx.NotFound()
	}

	field := app.SdqaliField {
		ID: ctx.FieldID,
		Name: fmt.Sprintf("Field %s", ctx.FieldID),
		Href: app.FieldHref(ctx.FieldID),
		Address: fmt.Sprintf("Address for %s", ctx.FieldID),
	}

	return ctx.OK(&field)
}
