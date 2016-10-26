package client

import (
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
func (c *Client) GetTransactionsTransactions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetTransactionsTransactionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetTransactionsTransactionsRequest create the request corresponding to the GetTransactions action endpoint of the transactions resource.
func (c *Client) NewGetTransactionsTransactionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
