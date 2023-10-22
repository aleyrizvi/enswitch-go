package enswitch

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

const (
	testUsername = "test_username"
	testPassword = "test_password"
	testBaseURL  = "test.test.com"
)

func TestClientNewRequest(t *testing.T) {
	c := Client{
		httpClient: http.DefaultClient,
		username:   testUsername,
		password:   testPassword,
		baseURL:    testBaseURL,
		Customer:   nil,
	}

	req, err := c.newRequest(context.Background(), "/test", url.Values{})
	if err != nil {
		t.Fatal("error occured")
	}

	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		t.Fatal("error occured")
	}

	if q.Get("auth_username") != testUsername || q.Get("auth_password") != testPassword {
		t.Error("error occured with query params")
	}
}

func TestClientCall(t *testing.T) {
}
