package enswitch

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	testUsername = "test_username"
	testPassword = "test_password"
	testBaseURL  = "test.test.com"
)

func TestClientNewRequest(t *testing.T) {
	type Test struct {
		expectedError error
	}

	cases := []Test{
		{
			expectedError: nil,
		},
	}

	for _, c := range cases {
		client := Client{
			httpClient: http.DefaultClient,
			username:   testUsername,
			password:   testPassword,
			baseURL:    testBaseURL,
			Customer:   nil,
		}

		req, err := client.newRequest(context.Background(), "/test", url.Values{})
		if !errors.Is(c.expectedError, err) {
			t.Errorf("error occured while creating new request. Given %v, Got %v", c.expectedError, err)
		}

		q, err := url.ParseQuery(req.URL.RawQuery)
		if !errors.Is(nil, err) {
			t.Errorf("error occured while parsing query params. Given %v, Got %v", c.expectedError, err)
		}

		if q.Get("auth_username") != testUsername || q.Get("auth_password") != testPassword {
			t.Error("error occured with query params")
		}
	}
}

func TestCall(t *testing.T) {
	type Test struct {
		body          string
		expectedError error
		expectedCode  uint
	}

	cases := []Test{
		{
			body:          `{"responses": [{"key": "","code": "204","message": "OK"}]}`,
			expectedError: nil,
			expectedCode:  204,
		},
	}

	for _, c := range cases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, c.body)
		}))
		defer server.Close()

		client := Client{
			httpClient: http.DefaultClient,
			username:   testUsername,
			password:   testPassword,
			baseURL:    server.URL,
			Customer:   nil,
		}

		req, err := client.newRequest(context.Background(), "/", url.Values{})
		if err != nil {
			t.Errorf("error occured while creating new request. Given %v, Got %v", c.expectedError, err)
		}

		var response BalanceResponse

		res, err := client.call(req, &response)

		if !errors.Is(err, c.expectedError) {
			t.Errorf("error occured calling http method. Given %v, Got %v", c.expectedError, err)
		}
		defer res.Body.Close()

		if response.Responses[0].Code != c.expectedCode {
			t.Errorf("status code: given: %v, got: %v", c.expectedCode, response.Responses[0].Code)
		}
	}
}
