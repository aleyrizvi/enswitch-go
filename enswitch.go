package enswitch

import (
	"fmt"
	"net/http"
)

// New returns Enswitch Client.
func New(username, password, baseUrl string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
		username:   username,
		password:   password,
		baseUrl:    fmt.Sprintf("https://%s/api/json/", baseUrl),
	}

	c.Customer = &Customer{
		client: c,
	}

	return c
}
