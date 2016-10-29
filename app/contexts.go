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
	Iban *string
}

// NewGetCashAccountsAccountsContext parses the incoming request URL and body, performs validations and creates the
// context used by the accounts controller GetCashAccounts action.
func NewGetCashAccountsAccountsContext(ctx context.Context, service *goa.Service) (*GetCashAccountsAccountsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetCashAccountsAccountsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramIban := req.Params["iban"]
	if len(paramIban) > 0 {
		rawIban := paramIban[0]
		rctx.Iban = &rawIban
	}
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

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GetCashAccountsAccountsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetCashAccountsAccountsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// GetTransactionsTransactionsContext provides the transactions GetTransactions action context.
type GetTransactionsTransactionsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Iban *string
}

// NewGetTransactionsTransactionsContext parses the incoming request URL and body, performs validations and creates the
// context used by the transactions controller GetTransactions action.
func NewGetTransactionsTransactionsContext(ctx context.Context, service *goa.Service) (*GetTransactionsTransactionsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetTransactionsTransactionsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramIban := req.Params["iban"]
	if len(paramIban) > 0 {
		rawIban := paramIban[0]
		rctx.Iban = &rawIban
	}
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

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GetTransactionsTransactionsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetTransactionsTransactionsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// PostTransactionsTransactionsContext provides the transactions PostTransactions action context.
type PostTransactionsTransactionsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PostTransactionsTransactionsPayload
}

// NewPostTransactionsTransactionsContext parses the incoming request URL and body, performs validations and creates the
// context used by the transactions controller PostTransactions action.
func NewPostTransactionsTransactionsContext(ctx context.Context, service *goa.Service) (*PostTransactionsTransactionsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := PostTransactionsTransactionsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// postTransactionsTransactionsPayload is the transactions PostTransactions action payload.
type postTransactionsTransactionsPayload struct {
	Amount                *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	CreditorBIC           *string `form:"creditorBIC,omitempty" json:"creditorBIC,omitempty" xml:"creditorBIC,omitempty"`
	CreditorIBAN          *string `form:"creditorIBAN,omitempty" json:"creditorIBAN,omitempty" xml:"creditorIBAN,omitempty"`
	Currency              *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	DebtorBIC             *string `form:"debtorBIC,omitempty" json:"debtorBIC,omitempty" xml:"debtorBIC,omitempty"`
	DebtorIBAN            *string `form:"debtorIBAN,omitempty" json:"debtorIBAN,omitempty" xml:"debtorIBAN,omitempty"`
	RemittanceInformation *string `form:"remittanceInformation,omitempty" json:"remittanceInformation,omitempty" xml:"remittanceInformation,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *postTransactionsTransactionsPayload) Validate() (err error) {
	if payload.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}
	if payload.CreditorIBAN == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "creditorIBAN"))
	}
	if payload.CreditorBIC == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "creditorBIC"))
	}
	if payload.DebtorIBAN == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "debtorIBAN"))
	}
	if payload.DebtorBIC == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "debtorBIC"))
	}
	if payload.Currency == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "currency"))
	}
	if payload.RemittanceInformation == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "remittanceInformation"))
	}

	return
}

// Publicize creates PostTransactionsTransactionsPayload from postTransactionsTransactionsPayload
func (payload *postTransactionsTransactionsPayload) Publicize() *PostTransactionsTransactionsPayload {
	var pub PostTransactionsTransactionsPayload
	if payload.Amount != nil {
		pub.Amount = *payload.Amount
	}
	if payload.CreditorBIC != nil {
		pub.CreditorBIC = *payload.CreditorBIC
	}
	if payload.CreditorIBAN != nil {
		pub.CreditorIBAN = *payload.CreditorIBAN
	}
	if payload.Currency != nil {
		pub.Currency = *payload.Currency
	}
	if payload.DebtorBIC != nil {
		pub.DebtorBIC = *payload.DebtorBIC
	}
	if payload.DebtorIBAN != nil {
		pub.DebtorIBAN = *payload.DebtorIBAN
	}
	if payload.RemittanceInformation != nil {
		pub.RemittanceInformation = *payload.RemittanceInformation
	}
	return &pub
}

// PostTransactionsTransactionsPayload is the transactions PostTransactions action payload.
type PostTransactionsTransactionsPayload struct {
	Amount                string `form:"amount" json:"amount" xml:"amount"`
	CreditorBIC           string `form:"creditorBIC" json:"creditorBIC" xml:"creditorBIC"`
	CreditorIBAN          string `form:"creditorIBAN" json:"creditorIBAN" xml:"creditorIBAN"`
	Currency              string `form:"currency" json:"currency" xml:"currency"`
	DebtorBIC             string `form:"debtorBIC" json:"debtorBIC" xml:"debtorBIC"`
	DebtorIBAN            string `form:"debtorIBAN" json:"debtorIBAN" xml:"debtorIBAN"`
	RemittanceInformation string `form:"remittanceInformation" json:"remittanceInformation" xml:"remittanceInformation"`
}

// Validate runs the validation rules defined in the design.
func (payload *PostTransactionsTransactionsPayload) Validate() (err error) {
	if payload.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}
	if payload.CreditorIBAN == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "creditorIBAN"))
	}
	if payload.CreditorBIC == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "creditorBIC"))
	}
	if payload.DebtorIBAN == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "debtorIBAN"))
	}
	if payload.DebtorBIC == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "debtorBIC"))
	}
	if payload.Currency == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "currency"))
	}
	if payload.RemittanceInformation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "remittanceInformation"))
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *PostTransactionsTransactionsContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PostTransactionsTransactionsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *PostTransactionsTransactionsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *PostTransactionsTransactionsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
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

// Created sends a HTTP response with status code 201.
func (ctx *GetAddressesUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
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
