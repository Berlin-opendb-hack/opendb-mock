package accounting

import (
	"sync"
	"github.com/shopspring/decimal"
	"time"
	"fmt"
)

var ledgerOnce sync.Once

type Ledger struct {
	Accounts map[string]*Account
}

var AccountingLedger Ledger

func GetLedger() *Ledger {
	ledgerOnce.Do(func() {
		AccountingLedger = Ledger{Accounts: make(map[string]*Account)}
	})
	return &AccountingLedger
}

func (l *Ledger) FindOrCreateAccount(identifier string) *Account {

	account := l.FindAccount(identifier)
	if (nil != account) {
		fmt.Println("found " + account.Identifier)
		return account
	}
	account = &Account{
		Identifier: identifier,
		credits:    []Credit{},
		debits :    []Debit{},
	}
	fmt.Println("created " + account.Identifier)
	l.Accounts[identifier] = account
	return account
}
func (l *Ledger) FindAccount(identifier string) *Account {

	account, ok := l.Accounts[identifier]
	if (!ok) {
		return nil
	}
	fmt.Println("found " + account.Identifier)
	return account
}

func (l *Ledger) Book(creditSideIdentifier string, debitSideIdentifier string, amount string, currency string, remittanceInformation string, date string) error {
	decAmount, err := decimal.NewFromString(amount)
	if (nil != err) {
		return err
	}
	debit := Debit{
		Amount: decAmount,
		Currency: currency,
		RemittanceInformation: remittanceInformation,
		Date: date,
		CreatedAt: time.Now(),
	}

	credit := Credit{
		Amount: decAmount,
		Currency: currency,
		RemittanceInformation:remittanceInformation,
		Date: date,
		CreatedAt: time.Now(),
	}
	creditSide := l.FindOrCreateAccount(creditSideIdentifier)
	debitSide := l.FindOrCreateAccount(debitSideIdentifier)
	creditSide.credits = append(creditSide.credits, credit)
	debitSide.debits = append(debitSide.debits, debit)


	return nil

}