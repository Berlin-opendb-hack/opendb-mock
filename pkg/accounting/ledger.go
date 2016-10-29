package accounting

import (
	"sync"
)

var ledgerOnce sync.Once

type Ledger struct {
	Accounts []Account
}

var AccountingLedger Ledger

func GetLedger() *Ledger {
	ledgerOnce.Do(func() {
		AccountingLedger = Ledger{Accounts: []Account{}}
	})
	return &AccountingLedger
}

func (l *Ledger) CreateAccount(identifier string) *Account {

	account := Account{
		Identifier: identifier,
		credits:    []Credit{},
		debits :    []Debit{},
	}
	l.Accounts = append(l.Accounts, account)
	return &account
}
func (l *Ledger) FindAccount(identifier string) *Account {

	for _, account := range l.Accounts {
		if account.Identifier == identifier {
			return &account
		}
	}
	return nil
}