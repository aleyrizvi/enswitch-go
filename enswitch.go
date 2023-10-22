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
func New(username, password, baseURL string) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeOut,
		},
		username: username,
		password: password,
		baseURL:  fmt.Sprintf("https://%s/api/json/", baseURL),
	}

	balance := &Balance{client: c}

	c.Customer = &Customer{client: c, Balance: balance}

	return c
}
