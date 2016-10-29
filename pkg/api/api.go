package api

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"encoding/json"
)

type CashAccount struct {
	Iban  string `json:"iban"`
	Balance float64 `json:"balance"`
	ProductDescription  string `json:"productDescription"`
}
type CashTransaction struct {
	Amount float64 `json:"amount"`
	CounterPartyName *string `json:"counterPartyName"`
	CounterPartyIban *string `json:"counterPartyIban"`
	Usage *string `json:"usage"`
	Date string `json:"date"`
}

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}

type DeutscheBankClient struct {
	httpClient Requestor
	endpoint url.URL
}


func NewDeutscheBankClient(endpoint url.URL, client Requestor) (*DeutscheBankClient, error) {
	return &DeutscheBankClient{httpClient: client, endpoint:endpoint}, nil
}

func (db *DeutscheBankClient) GetCashAccounts (token string, iban string) ([]CashAccount, error) {
	endpoint := url.URL{
		Scheme: db.endpoint.Scheme,
		Host: db.endpoint.Host,
		Path: db.endpoint.Path + "/cashAccounts",

	}
	if ("" != iban) {
		endpoint.RawQuery = fmt.Sprintf("iban=%s", iban)
	}
	request, err := http.NewRequest("GET",endpoint.String(), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	if (nil != err) {
		return []CashAccount{}, err
	}

	res, err := db.httpClient.Do(request)
	if (nil != err) {
		return []CashAccount{}, err
	}
	if (res.StatusCode == 200) {
		cashAccounts := []CashAccount{}
		responseBody, err := ioutil.ReadAll(res.Body)
		if (nil != err ) {
			return []CashAccount{}, err
		}
		json.Unmarshal(responseBody, &cashAccounts)

		return cashAccounts, nil
	}
	return []CashAccount{}, errors.Errorf("Unsupported status code %s", res.Status)

}

func (db *DeutscheBankClient) GetTransactions (token string, iban string) ([]CashTransaction, error) {
	endpoint := url.URL{
		Scheme: db.endpoint.Scheme,
		Host: db.endpoint.Host,
		Path: db.endpoint.Path + "/transactions",

	}
	if ("" != iban) {
		endpoint.RawQuery = fmt.Sprintf("iban=%s", iban)
	}
	request, err := http.NewRequest("GET",endpoint.String(), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	if (nil != err) {
		return []CashTransaction{}, err
	}

	res, err := db.httpClient.Do(request)
	if (res.StatusCode == 200) {
		txs := []CashTransaction{}
		responseBody, err := ioutil.ReadAll(res.Body)
		if (nil != err ) {
			return []CashTransaction{}, err
		}
		json.Unmarshal(responseBody, &txs)

		return txs, nil
	}
	return []CashTransaction{}, errors.Errorf("Unsupported status code %s", res.Status)
}
