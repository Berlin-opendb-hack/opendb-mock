package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("opendb", func() {
	// API defines the microservice endpoint and
	Title("The mock for opendb")           // other global properties. There should be one
	Description("Service to proxy ")        // and exactly one API definition appearing in
	Scheme("http")                             // the design.
	Host("localhost:8880")
	APIKeySecurity("token", func() {
		Header("token")
	})

})

var _ = Resource("transactions", func() {
	BasePath("/transactions")
	Security("token", func() {})
	Action("GetTransactions", func() {
		Description("Get all transactions")
		Routing(GET("/"))
		Params(func() {
			Param("iban", String, func() {

			})
		})
		Response(OK, func() {
			Media(CollectionOf(TransactionMedia))
		})
		Response(Unauthorized)
		Response(InternalServerError, func() {
			Media(ErrorMedia)
		})


	})

	Action("PostTransactions", func() {
		Description("Create a transaction")
		Routing(POST("/"))
		Response(Created)
		Response(Unauthorized)
		Response(InternalServerError, func() {
			Media(ErrorMedia)
		})
		Payload(SepaTransferMedia)
	})
})
var _ = Resource("accounts", func() {
	BasePath("/cashAccounts")
	Security("token", func() {})
	Action("GetCashAccounts", func() {
		Description("Cash accuonts")
		Routing(GET("/"))
		Params(func() {
			Param("iban", String, func() {

			})
		})
		Response(OK, func() {
			Media(CollectionOf(CashAccountMedia))
		})
		Response(Unauthorized)
		Response(InternalServerError, func() {
			Media(ErrorMedia)
		})
	})

})
var _ = Resource("user", func() {
	BasePath("/")

	Action("GetAddresses", func() {
		Description("Get adresses")
		Routing(GET("addresses"))
		Response(OK, func() {
			Media(CollectionOf(AddressMedia))
		})
		Response(Created)
	})

	Action("GetUserInfo", func() {
		Description("Get User Info")
		Routing(GET("userInfo"))
		Response(OK, func() {
			Media(UserInfoMedia)
		})
	})
})

var SepaTransferMedia = MediaType("application/vnd.opendb.hack.sepa+json", func() {
	Description("")
	Attributes(func() {
		Attribute("amount", String, "")
		Attribute("creditorIBAN", String, "")
		Attribute("creditorBIC", String, "")
		Attribute("debtorIBAN", String, "")
		Attribute("debtorBIC", String, "")
		Attribute("currency", String, "")
		Attribute("remittanceInformation", String, "")
	})

	Required(
		"amount",
		"creditorIBAN",
		"creditorBIC",
		"debtorIBAN",
		"debtorBIC",
		"currency",
		"remittanceInformation",
	)
	View("default", func() {
		Attribute("amount")
		Attribute("currency")
		Attribute("remittanceInformation")
	})
})

var TransactionMedia = MediaType("application/vnd.opendb.hack.transaction+json", func() {
	Attributes(func() {
		Attribute("amount", Number, "Amount of the transaction. If the amount is positive, the customer received money, if the amount is negative the customer lost money,")
		Attribute("date", String, "Posting date in the format YYYY-MM-DD")
		Attribute("usage", String, "Payment reference")
		Attribute("counterPartyName", String, "Name of the counter party")
		Attribute("counterPartyIban", String, "IBAN of the counter party")
	})
	Required(
		"amount",
		"date",
	)
	View("default", func() {
		Attribute("amount")
		Attribute("date")
		Attribute("usage")
		Attribute("counterPartyName")
		Attribute("counterPartyIban")
	})
})

var UserInfoMedia = MediaType("application/vnd.opendb.hack.userinfo+json", func() {
	Attribute("firstName", String, "First name of the user")
	Attribute("lastName", String, "Last name of the user")
	Attribute("dateOfBirth", String, "Birth date of the user in the format YYYY-MM-DD")
	Attribute("gender", String, "Gender of the user = [„MALE“ or „FEMALE“ or „UNKOWN“]")
	Required(
		"firstName",
		"lastName",
		"dateOfBirth",
		"gender",
	)
	View("default", func() {
		Attribute("firstName")
		Attribute("lastName")
		Attribute("dateOfBirth")
		Attribute("gender")
	})
})

var AddressMedia = MediaType("application/vnd.opendb.hack.address+json", func() {
	Attribute("street", String, "Street,")
	Attribute("houseNumber", String, "House number,")
	Attribute("zip", String, "Zip code,")
	Attribute("city", String, "City,")
	Attribute("country", String, "Country Code,")
	Attribute("type", String, "Type of address = [„MAILING_ADDRESS“ or „REGISTRATION_ADDRESS“]")
	Required(
		"street",
		"houseNumber",
		"zip",
		"city",
		"country",
		"type",
	)
	View("default", func() {
		Attribute("street")
		Attribute("houseNumber")
		Attribute("zip")
		Attribute("city")
		Attribute("country")
		Attribute("type")

	})
})

var CashAccountMedia = MediaType("application/vnd.opendb.hack.cashaccount+json", func() {
	Attribute("iban", String, "IBAN of the cash account")
	Attribute("balance", Number, "Booked balance in EUR")
	Attribute("productDescription", String, "Description of the product")
	Required(
		"iban",
		"balance",
		"productDescription",
	)
	View("default", func() {
		Attribute("iban")
		Attribute("balance")
		Attribute("productDescription")
	})
})