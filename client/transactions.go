package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetTransactionsTransactionsPath computes a request path to the GetTransactions action of transactions.
func GetTransactionsTransactionsPath() string {
	return fmt.Sprintf("/transactions")
}

// Get all transactions
func (c *Client) GetTransactionsTransactions(ctx context.Context, path string, iban *string) (*http.Response, error) {
	req, err := c.NewGetTransactionsTransactionsRequest(ctx, path, iban)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetTransactionsTransactionsRequest create the request corresponding to the GetTransactions action endpoint of the transactions resource.
func (c *Client) NewGetTransactionsTransactionsRequest(ctx context.Context, path string, iban *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if iban != nil {
		values.Set("iban", *iban)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.TokenSigner != nil {
		c.TokenSigner.Sign(req)
	}
	return req, nil
}

// PostTransactionsTransactionsPayload is the transactions PostTransactions action payload.
type PostTransactionsTransactionsPayload struct {
	Amount                string `form:"amount" json:"amount" xml:"amount"`
	CreditorBIC           string `form:"creditorBIC" json:"creditorBIC" xml:"creditorBIC"`
	CreditorIBAN          string `form:"creditorIBAN" json:"creditorIBAN" xml:"creditorIBAN"`
	Currency              string `form:"currency" json:"currency" xml:"currency"`
	DebtorBIC             string `form:"debtorBIC" json:"debtorBIC" xml:"debtorBIC"`
	DebtorIBAN            string `form:"debtorIBAN" json:"debtorIBAN" xml:"debtorIBAN"`
	RemittanceInformation string `form:"remittanceInformation" json:"remittanceInformation" xml:"remittanceInformation"`
}

// PostTransactionsTransactionsPath computes a request path to the PostTransactions action of transactions.
func PostTransactionsTransactionsPath() string {
	return fmt.Sprintf("/transactions")
}

// Create a transaction
func (c *Client) PostTransactionsTransactions(ctx context.Context, path string, payload *PostTransactionsTransactionsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewPostTransactionsTransactionsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPostTransactionsTransactionsRequest create the request corresponding to the PostTransactions action endpoint of the transactions resource.
func (c *Client) NewPostTransactionsTransactionsRequest(ctx context.Context, path string, payload *PostTransactionsTransactionsPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	if c.TokenSigner != nil {
		c.TokenSigner.Sign(req)
	}
	return req, nil
}
