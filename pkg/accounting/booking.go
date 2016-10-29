package accounting

import (
	"github.com/shopspring/decimal"
	"time"
)

func Book(creditSide *Account, debitSide *Account, amount string, currency string, remittanceInformation string, date string) error {
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
	creditSide.credits = append(creditSide.credits, credit)
	debitSide.debits = append(debitSide.debits, debit)

	return nil

}
