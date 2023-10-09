package enswitch

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultTimeOut = time.Minute
)

// New returns Enswitch Client.
func New(username, password, baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeOut,
		}
	}

	c := &Client{
		httpClient: httpClient,
		username:   username,
		password:   password,
		baseURL:    fmt.Sprintf("https://%s/api/json/", baseURL),
	}

	c.Customer = &Customer{
		client: c,
	}

	return c
}
