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
	username, password, baseUrl string
	Customer                    *Customer
}

func (c *Client) newRequest(ctx context.Context, uri string, qs url.Values) (*http.Request, error) {
	u, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, ErrBadUrl
	}

	u.Path += uri

	if ctx == nil {
		ctx = context.Background()
	}

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

	if !(res.StatusCode > 200) && !(res.StatusCode < 300) {
		return nil, fmt.Errorf("server responded with error code: %d", res.StatusCode)
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
