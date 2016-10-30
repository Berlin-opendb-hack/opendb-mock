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
	"github.com/goadesign/goa/cors"
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
	service.Mux.Handle("OPTIONS", "/cashAccounts", ctrl.MuxHandler("preflight", handleAccountsOrigin(cors.HandlePreflight()), nil))

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
	h = handleAccountsOrigin(h)
	h = handleSecurity("token", h)
	service.Mux.Handle("GET", "/cashAccounts", ctrl.MuxHandler("GetCashAccounts", h, nil))
	service.LogInfo("mount", "ctrl", "Accounts", "action", "GetCashAccounts", "route", "GET /cashAccounts", "security", "token")
}

// handleAccountsOrigin applies the CORS response headers corresponding to the origin.
func handleAccountsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8089") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/transactions", ctrl.MuxHandler("preflight", handleTransactionsOrigin(cors.HandlePreflight()), nil))

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
	h = handleTransactionsOrigin(h)
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
	h = handleTransactionsOrigin(h)
	h = handleSecurity("token", h)
	service.Mux.Handle("POST", "/transactions", ctrl.MuxHandler("PostTransactions", h, unmarshalPostTransactionsTransactionsPayload))
	service.LogInfo("mount", "ctrl", "Transactions", "action", "PostTransactions", "route", "POST /transactions", "security", "token")
}

// handleTransactionsOrigin applies the CORS response headers corresponding to the origin.
func handleTransactionsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8089") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/addresses", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/userInfo", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

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
	h = handleUserOrigin(h)
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
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/userInfo", ctrl.MuxHandler("GetUserInfo", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GetUserInfo", "route", "GET /userInfo")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8089") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "X-Token")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
