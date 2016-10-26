package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetCashAccountsAccountsPath computes a request path to the GetCashAccounts action of accounts.
func GetCashAccountsAccountsPath() string {
	return fmt.Sprintf("/cashAccounts")
}

// Cash accuonts
func (c *Client) GetCashAccountsAccounts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCashAccountsAccountsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCashAccountsAccountsRequest create the request corresponding to the GetCashAccounts action endpoint of the accounts resource.
func (c *Client) NewGetCashAccountsAccountsRequest(ctx context.Context, path string) (*http.Request, error) {
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
