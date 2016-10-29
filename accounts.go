package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/api"
	"net/http"
	"net/url"
	"os"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/accounting"
	"github.com/shopspring/decimal"
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
	res := app.OpendbHackCashaccountCollection{}
	client := &http.Client{}
	gateWay, err := api.NewDeutscheBankClient(
		url.URL{Scheme: os.Getenv("OPENDB_SCHEME"), Host: os.Getenv("OPENDB_HOST"), Path: os.Getenv("OPENDB_PATH")},
		client,
	)
	if (nil != err) {
		return ctx.InternalServerError(err)
	}
	ledger := accounting.GetLedger()
	token := ctx.Value("token")
	if nil == token {
		return ctx.Unauthorized()
	}
	cashAccounts, err := gateWay.GetCashAccounts(token.(string), ctx.Params.Get("iban"))
	if (nil != err) {
		return ctx.InternalServerError(err)
	}

	for _, cashAccount := range cashAccounts {
		account := ledger.FindAccount(cashAccount.Iban)
		if nil == account {
			account = ledger.CreateAccount(cashAccount.Iban)
		}
		realBalance := decimal.NewFromFloat(cashAccount.Balance)
		realBalance = realBalance.Add(account.Balance())
		balance, _ := realBalance.Float64()
		resAccount := app.OpendbHackCashaccount{
			Iban: cashAccount.Iban,
			Balance: balance,
			ProductDescription: cashAccount.ProductDescription,
		}
		res = append(res, &resAccount)
	}

	return ctx.OK(res)
}
