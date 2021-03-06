//************************************************************************//
// API "opendb": Application Media Types
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

import "github.com/goadesign/goa"

// OpendbHackAddress media type (default view)
//
// Identifier: application/vnd.opendb.hack.address+json; view=default
type OpendbHackAddress struct {
	// City,
	City string `form:"city" json:"city" xml:"city"`
	// Country Code,
	Country string `form:"country" json:"country" xml:"country"`
	// House number,
	HouseNumber string `form:"houseNumber" json:"houseNumber" xml:"houseNumber"`
	// Street,
	Street string `form:"street" json:"street" xml:"street"`
	// Type of address = [„MAILING_ADDRESS“ or „REGISTRATION_ADDRESS“]
	Type string `form:"type" json:"type" xml:"type"`
	// Zip code,
	Zip string `form:"zip" json:"zip" xml:"zip"`
}

// Validate validates the OpendbHackAddress media type instance.
func (mt *OpendbHackAddress) Validate() (err error) {
	if mt.Street == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street"))
	}
	if mt.HouseNumber == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "houseNumber"))
	}
	if mt.Zip == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}
	if mt.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if mt.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// OpendbHackAddressCollection is the media type for an array of OpendbHackAddress (default view)
//
// Identifier: application/vnd.opendb.hack.address+json; type=collection; view=default
type OpendbHackAddressCollection []*OpendbHackAddress

// Validate validates the OpendbHackAddressCollection media type instance.
func (mt OpendbHackAddressCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Street == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "street"))
		}
		if e.HouseNumber == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "houseNumber"))
		}
		if e.Zip == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "zip"))
		}
		if e.City == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "city"))
		}
		if e.Country == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "country"))
		}
		if e.Type == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "type"))
		}

	}
	return
}

// OpendbHackCashaccount media type (default view)
//
// Identifier: application/vnd.opendb.hack.cashaccount+json; view=default
type OpendbHackCashaccount struct {
	// Booked balance in EUR
	Balance float64 `form:"balance" json:"balance" xml:"balance"`
	// IBAN of the cash account
	Iban string `form:"iban" json:"iban" xml:"iban"`
	// Description of the product
	ProductDescription string `form:"productDescription" json:"productDescription" xml:"productDescription"`
}

// Validate validates the OpendbHackCashaccount media type instance.
func (mt *OpendbHackCashaccount) Validate() (err error) {
	if mt.Iban == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "iban"))
	}
	if mt.ProductDescription == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "productDescription"))
	}

	return
}

// OpendbHackCashaccountCollection is the media type for an array of OpendbHackCashaccount (default view)
//
// Identifier: application/vnd.opendb.hack.cashaccount+json; type=collection; view=default
type OpendbHackCashaccountCollection []*OpendbHackCashaccount

// Validate validates the OpendbHackCashaccountCollection media type instance.
func (mt OpendbHackCashaccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Iban == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "iban"))
		}
		if e.ProductDescription == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "productDescription"))
		}

	}
	return
}

// OpendbHackSepa media type (default view)
//
// Identifier: application/vnd.opendb.hack.sepa+json; view=default
type OpendbHackSepa struct {
	Amount                string `form:"amount" json:"amount" xml:"amount"`
	Currency              string `form:"currency" json:"currency" xml:"currency"`
	RemittanceInformation string `form:"remittanceInformation" json:"remittanceInformation" xml:"remittanceInformation"`
}

// Validate validates the OpendbHackSepa media type instance.
func (mt *OpendbHackSepa) Validate() (err error) {
	if mt.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}
	if mt.Currency == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "currency"))
	}
	if mt.RemittanceInformation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "remittanceInformation"))
	}

	return
}

// OpendbHackTransaction media type (default view)
//
// Identifier: application/vnd.opendb.hack.transaction+json; view=default
type OpendbHackTransaction struct {
	// Amount of the transaction. If the amount is positive, the customer received money, if the amount is negative the customer lost money,
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// IBAN of the counter party
	CounterPartyIban *string `form:"counterPartyIban,omitempty" json:"counterPartyIban,omitempty" xml:"counterPartyIban,omitempty"`
	// Name of the counter party
	CounterPartyName *string `form:"counterPartyName,omitempty" json:"counterPartyName,omitempty" xml:"counterPartyName,omitempty"`
	// Posting date in the format YYYY-MM-DD
	Date string `form:"date" json:"date" xml:"date"`
	// Payment reference
	Usage *string `form:"usage,omitempty" json:"usage,omitempty" xml:"usage,omitempty"`
}

// Validate validates the OpendbHackTransaction media type instance.
func (mt *OpendbHackTransaction) Validate() (err error) {
	if mt.Date == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "date"))
	}

	return
}

// OpendbHackTransactionCollection is the media type for an array of OpendbHackTransaction (default view)
//
// Identifier: application/vnd.opendb.hack.transaction+json; type=collection; view=default
type OpendbHackTransactionCollection []*OpendbHackTransaction

// Validate validates the OpendbHackTransactionCollection media type instance.
func (mt OpendbHackTransactionCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Date == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "date"))
		}

	}
	return
}

// OpendbHackUserinfo media type (default view)
//
// Identifier: application/vnd.opendb.hack.userinfo+json; view=default
type OpendbHackUserinfo struct {
	// Birth date of the user in the format YYYY-MM-DD
	DateOfBirth string `form:"dateOfBirth" json:"dateOfBirth" xml:"dateOfBirth"`
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Gender of the user = [„MALE“ or „FEMALE“ or „UNKOWN“]
	Gender string `form:"gender" json:"gender" xml:"gender"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
}

// Validate validates the OpendbHackUserinfo media type instance.
func (mt *OpendbHackUserinfo) Validate() (err error) {
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "lastName"))
	}
	if mt.DateOfBirth == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "dateOfBirth"))
	}
	if mt.Gender == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "gender"))
	}

	return
}
