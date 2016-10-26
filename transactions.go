package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
)

// TransactionsController implements the transactions resource.
type TransactionsController struct {
	*goa.Controller
}

// NewTransactionsController creates a transactions controller.
func NewTransactionsController(service *goa.Service) *TransactionsController {
	return &TransactionsController{Controller: service.NewController("TransactionsController")}
}

// GetTransactions runs the GetTransactions action.
func (c *TransactionsController) GetTransactions(ctx *app.GetTransactionsTransactionsContext) error {
	// TransactionsController_GetTransactions: start_implement

	// Put your logic here

	// TransactionsController_GetTransactions: end_implement
	res := app.OpendbHackTransactionCollection{}
	return ctx.OK(res)
}
