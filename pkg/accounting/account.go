package accounting

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	Amount                decimal.Decimal
	CounterPartyIban      *string
	CounterPartyName      *string
	Currency              string
	RemittanceInformation string
	Date                  string
	CreatedAt             time.Time
}
type Credit Transaction
type Debit Transaction
type Account struct {
	Identifier string
	credits    []Credit
	debits     []Debit
}

func (a Account) Balance() decimal.Decimal {
	total := decimal.Zero
	for _, credit := range a.credits {
		total = total.Add(credit.Amount)
	}
	for _, debit := range a.debits {
		total = total.Sub(debit.Amount)
	}

	return total
}

func (a Account) GetCredits() []Credit {
	return a.credits
}
func (a Account) GetDebits() []Debit {
	return a.debits
}

