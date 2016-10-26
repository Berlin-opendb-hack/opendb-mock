package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
)

// AccountsController implements the accounts resource.
type AccountsController struct {
	*goa.Controller
}

// NewAccountsController creates a accounts controller.
func NewAccountsController(service *goa.Service) *AccountsController {
	return &AccountsController{Controller: service.NewController("AccountsController")}
}

// GetCashAccounts runs the GetCashAccounts action.
func (c *AccountsController) GetCashAccounts(ctx *app.GetCashAccountsAccountsContext) error {
	// AccountsController_GetCashAccounts: start_implement

	// Put your logic here

	// AccountsController_GetCashAccounts: end_implement
	res := app.OpendbHackCashaccountCollection{}
	return ctx.OK(res)
}
