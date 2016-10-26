//************************************************************************//
// API "opendb": Application Contexts
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
)

// GetCashAccountsAccountsContext provides the accounts GetCashAccounts action context.
type GetCashAccountsAccountsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetCashAccountsAccountsContext parses the incoming request URL and body, performs validations and creates the
// context used by the accounts controller GetCashAccounts action.
func NewGetCashAccountsAccountsContext(ctx context.Context, service *goa.Service) (*GetCashAccountsAccountsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetCashAccountsAccountsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetCashAccountsAccountsContext) OK(r OpendbHackCashaccountCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.cashaccount+json; type=collection")
	if r == nil {
		r = OpendbHackCashaccountCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// GetTransactionsTransactionsContext provides the transactions GetTransactions action context.
type GetTransactionsTransactionsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetTransactionsTransactionsContext parses the incoming request URL and body, performs validations and creates the
// context used by the transactions controller GetTransactions action.
func NewGetTransactionsTransactionsContext(ctx context.Context, service *goa.Service) (*GetTransactionsTransactionsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetTransactionsTransactionsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetTransactionsTransactionsContext) OK(r OpendbHackTransactionCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.transaction+json; type=collection")
	if r == nil {
		r = OpendbHackTransactionCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// GetAddressesUserContext provides the user GetAddresses action context.
type GetAddressesUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetAddressesUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller GetAddresses action.
func NewGetAddressesUserContext(ctx context.Context, service *goa.Service) (*GetAddressesUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetAddressesUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetAddressesUserContext) OK(r OpendbHackAddressCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.address+json; type=collection")
	if r == nil {
		r = OpendbHackAddressCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// GetUserInfoUserContext provides the user GetUserInfo action context.
type GetUserInfoUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetUserInfoUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller GetUserInfo action.
func NewGetUserInfoUserContext(ctx context.Context, service *goa.Service) (*GetUserInfoUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetUserInfoUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUserInfoUserContext) OK(r *OpendbHackUserinfo) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.userinfo+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
