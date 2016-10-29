package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
	"net/url"
	"os"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/accounting"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/api"
	"net/http"
	"time"
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
	cashTransactions, err := gateWay.GetTransactions(token.(string), ctx.Params.Get("iban"))
	if (nil != err) {
		return ctx.InternalServerError(err)
	}
	res := app.OpendbHackTransactionCollection{}
	for _, tx := range cashTransactions {
		transaction := app.OpendbHackTransaction{
			Amount: tx.Amount,
			CounterPartyIban: tx.CounterPartyIban,
			CounterPartyName: tx.CounterPartyName,
			Date: tx.Date,
			Usage: tx.Usage,
		}
		res = append(res, &transaction)
	}
	for _, acc := range cashAccounts {
		account := ledger.FindAccount(acc.Iban)
		if (nil == account) {
			continue
		}
		for _, credit := range account.GetCredits() {
			amnt, _ := credit.Amount.Float64()
			usage := &credit.RemittanceInformation
			transaction := app.OpendbHackTransaction{
				Amount: amnt,
				CounterPartyIban: credit.CounterPartyIban,
				CounterPartyName: credit.CounterPartyName,
				Date: credit.Date,
				Usage: usage,
			}
			res = append(res, &transaction)
		}

		for _, debit := range account.GetCredits() {
			amnt, _ := debit.Amount.Float64()
			usage := &debit.RemittanceInformation
			transaction := app.OpendbHackTransaction{
				Amount: -amnt,
				CounterPartyIban: debit.CounterPartyIban,
				CounterPartyName: debit.CounterPartyName,
				Date: debit.Date,
				Usage: usage,
			}
			res = append(res, &transaction)
		}
	}
	return ctx.OK(res)
}

// PostTransactions runs the PostTransactions action.
func (c *TransactionsController) PostTransactions(ctx *app.PostTransactionsTransactionsContext) error {
	token := ctx.Value("token")
	if nil == token {
		return ctx.Unauthorized()
	}
	ledger := accounting.GetLedger()
	sepaTransfer := ctx.Payload
	validationErr := sepaTransfer.Validate()
	if (nil != validationErr) {
		return ctx.BadRequest(validationErr)
	}
	err := ledger.Book(sepaTransfer.CreditorIBAN, sepaTransfer.DebtorIBAN, sepaTransfer.Amount, sepaTransfer.Currency, sepaTransfer.RemittanceInformation, time.Now().Format("2006-01-02"))
	if (nil != err) {
		return ctx.InternalServerError(err)
	}
	return ctx.Created()
}
