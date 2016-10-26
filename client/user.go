package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetAddressesUserPath computes a request path to the GetAddresses action of user.
func GetAddressesUserPath() string {
	return fmt.Sprintf("/addresses")
}

// Get adresses
func (c *Client) GetAddressesUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetAddressesUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAddressesUserRequest create the request corresponding to the GetAddresses action endpoint of the user resource.
func (c *Client) NewGetAddressesUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetUserInfoUserPath computes a request path to the GetUserInfo action of user.
func GetUserInfoUserPath() string {
	return fmt.Sprintf("/userInfo")
}

// Get User Info
func (c *Client) GetUserInfoUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserInfoUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserInfoUserRequest create the request corresponding to the GetUserInfo action endpoint of the user resource.
func (c *Client) NewGetUserInfoUserRequest(ctx context.Context, path string) (*http.Request, error) {
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
