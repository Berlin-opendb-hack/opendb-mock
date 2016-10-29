//************************************************************************//
// API "opendb": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/Berlin-opendb-hack/opendb-mock/design
// --out=$(GOPATH)/src/github.com/Berlin-opendb-hack/opendb-mock
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountsController is the controller interface for the Accounts actions.
type AccountsController interface {
	goa.Muxer
	GetCashAccounts(*GetCashAccountsAccountsContext) error
}

// MountAccountsController "mounts" a Accounts resource controller on the given service.
func MountAccountsController(service *goa.Service, ctrl AccountsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetCashAccountsAccountsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetCashAccounts(rctx)
	}
	h = handleSecurity("token", h)
	service.Mux.Handle("GET", "/cashAccounts", ctrl.MuxHandler("GetCashAccounts", h, nil))
	service.LogInfo("mount", "ctrl", "Accounts", "action", "GetCashAccounts", "route", "GET /cashAccounts", "security", "token")
}

// TransactionsController is the controller interface for the Transactions actions.
type TransactionsController interface {
	goa.Muxer
	GetTransactions(*GetTransactionsTransactionsContext) error
	PostTransactions(*PostTransactionsTransactionsContext) error
}

// MountTransactionsController "mounts" a Transactions resource controller on the given service.
func MountTransactionsController(service *goa.Service, ctrl TransactionsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetTransactionsTransactionsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetTransactions(rctx)
	}
	h = handleSecurity("token", h)
	service.Mux.Handle("GET", "/transactions", ctrl.MuxHandler("GetTransactions", h, nil))
	service.LogInfo("mount", "ctrl", "Transactions", "action", "GetTransactions", "route", "GET /transactions", "security", "token")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPostTransactionsTransactionsContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostTransactionsTransactionsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.PostTransactions(rctx)
	}
	h = handleSecurity("token", h)
	service.Mux.Handle("POST", "/transactions", ctrl.MuxHandler("PostTransactions", h, unmarshalPostTransactionsTransactionsPayload))
	service.LogInfo("mount", "ctrl", "Transactions", "action", "PostTransactions", "route", "POST /transactions", "security", "token")
}

// unmarshalPostTransactionsTransactionsPayload unmarshals the request body into the context request data Payload field.
func unmarshalPostTransactionsTransactionsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postTransactionsTransactionsPayload{}
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
	GetAddresses(*GetAddressesUserContext) error
	GetUserInfo(*GetUserInfoUserContext) error
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
		rctx, err := NewGetAddressesUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetAddresses(rctx)
	}
	service.Mux.Handle("GET", "/addresses", ctrl.MuxHandler("GetAddresses", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GetAddresses", "route", "GET /addresses")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserInfoUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetUserInfo(rctx)
	}
	service.Mux.Handle("GET", "/userInfo", ctrl.MuxHandler("GetUserInfo", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GetUserInfo", "route", "GET /userInfo")
}
