//go:generate goagen bootstrap -d github.com/Berlin-opendb-hack/opendb-mock/design

package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("opendb")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "accounts" controller
	c := NewAccountsController(service)
	app.MountAccountsController(service, c)
	// Mount "transactions" controller
	c2 := NewTransactionsController(service)
	app.MountTransactionsController(service, c2)
	// Mount "user" controller
	c3 := NewUserController(service)
	app.MountUserController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8880"); err != nil {
		service.LogError("startup", "err", err)
	}
}
