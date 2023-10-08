package enswitch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeOut = time.Minute
)

type Backend struct {
	username string
	password string
	baseUrl  string
	client   *http.Client
}

func (b *Backend) Call(ctx context.Context, uri string, data map[string]string) (int, []byte, error) {
	u := fmt.Sprintf("%s?%s", b.getBaseURL(uri), b.getQueryParams(data).Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return 0, nil, ErrRequestWithContext
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("unable to perform http request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0, nil, fmt.Errorf("unable to read malformed server response")
	}

	return resp.StatusCode, body, nil

}

func (b *Backend) getBaseURL(uri string) string {
	return fmt.Sprintf("https://%s/api/json/%s", b.baseUrl, uri)
}

func (b *Backend) getQueryParams(data map[string]string) url.Values {
	q := url.Values{}

	q.Add("auth_username", b.username)
	q.Add("auth_password", b.password)

	for k, v := range data {
		q.Add(k, v)
	}

	return q
}
