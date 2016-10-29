//go:generate goagen bootstrap -d github.com/Berlin-opendb-hack/opendb-mock/design

package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"net/http"
	"golang.org/x/net/context"
)

func main() {
	// Create service
	service := goa.New("opendb")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseTokenMiddleware(service, AuthMiddleWare)

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

func AuthMiddleWare(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		token := req.Header.Get("X-Token")
		if ("" == token) {
			return goa.ErrUnauthorized("Token header required")
		}
		newctx := context.WithValue(ctx, "token", token)

		return h(newctx, rw, req)

	}
}

