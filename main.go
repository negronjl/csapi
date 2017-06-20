//go:generate goagen bootstrap -d github.com/negronjl/csapi/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/negronjl/csapi/app"
	"github.com/negronjl/csapi/controllers"
	"fmt"
	"os"
)

func main() {
	// Process environment
	zapi_list, err := controllers.Init_env()
	if err != nil {
		fmt.Printf("Unable to initialize environment. %v\n", err)
		os.Exit(1)
	}

	// Create service
	service := goa.New("CSAPI")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "health" controller
	c := controllers.NewHealthController(service, zapi_list)
	app.MountHealthController(service, c)
	// Mount "swagger" controller
	c2 := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
