package enswitch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient                  *http.Client
	username, password, baseURL string
	Customer                    *Customer
}

func (c *Client) newRequest(ctx context.Context, uri string, qs url.Values) (*http.Request, error) {
	if ctx == nil {
		return nil, ErrContextNil
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, ErrBadURL
	}

	u.Path += uri

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, ErrRequestWithContext
	}

	req.Header.Add("Content-Type", "Application/JSON")
	req.Header.Add("User-Agent", "Enswitch-GO")

	req.URL.RawQuery = c.parseQueryParams(qs).Encode()

	return req, nil
}

func (c *Client) call(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, ErrHTTPRequest
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code: %w: %d", ErrHTTPRequest, res.StatusCode)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)

	if err != nil {
		return nil, ErrDecodingRequest
	}

	return res, nil
}

func (c *Client) parseQueryParams(qs url.Values) url.Values {
	qs.Add("auth_username", c.username)
	qs.Add("auth_password", c.password)

	return qs
}
