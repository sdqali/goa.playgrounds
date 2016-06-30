package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"playgrounds/app"
)

func main() {
	// Create service
	service := goa.New("playgrounds")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "field" controller
	c := NewFieldController(service)
	app.MountFieldController(service, c)

	// Start service
	if err := service.ListenAndServe(":9876"); err != nil {
		service.LogError("startup", "err", err)
	}
}
