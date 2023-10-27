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

func TestBalance_Update(t *testing.T) {
	type Test struct {
		server        *httptest.Server
		input         BalanceParams
		errorExpected error
	}

	cases := []Test{
		{
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"responses": [{"key": "","code": "204","message": "OK"}]}`)
			})),
			input: BalanceParams{
				ID:     1,
				Amount: 10,
			},
			errorExpected: nil,
		},
	}

	for _, c := range cases {
		client := Client{
			httpClient: http.DefaultClient,
			username:   testUsername,
			password:   testPassword,
			baseURL:    c.server.URL,
			Customer:   nil,
		}

		balance := Balance{client: &client}

		res, err := balance.Update(context.Background(), c.input)

		if !errors.Is(err, c.errorExpected) {
			t.Errorf("[balanceUpdate]: given %v, got %v", c.errorExpected, err)
		}

		if res.Responses[0].Code != 204 {
			t.Error("[balanceUpdate]: status code does not match")
		}
	}
}

func TestCustomerBalanceQueryParams(t *testing.T) {
	transaction := TransactionTypeTopUp

	cases := []struct {
		params   BalanceParams
		expected url.Values
	}{
		{
			params: BalanceParams{
				ID:     123,
				Amount: 45.67,
			},
			expected: url.Values{
				"id":     []string{"123"},
				"amount": []string{"45.67"},
			},
		},
		{
			params: BalanceParams{
				ID:       456,
				Amount:   78.90,
				Currency: PointerTo("USD"),
			},
			expected: url.Values{
				"id":       []string{"456"},
				"amount":   []string{"78.9"},
				"currency": []string{"USD"},
			},
		},
		{
			params: BalanceParams{
				ID:          789,
				Amount:      12.34,
				Currency:    PointerTo("EUR"),
				Transaction: &transaction,
				Description: PointerTo("Test description"),
			},
			expected: url.Values{
				"id":          []string{"789"},
				"amount":      []string{"12.34"},
				"currency":    []string{"EUR"},
				"transaction": []string{"topup"},
				"description": []string{"Test description"},
			},
		},
	}

	for _, c := range cases {
		qs := balanceQueryParam(c.params)

		if !urlValuesEqual(qs, c.expected) {
			t.Errorf("[Balance] unable to create query param. Given %v, got %v", qs, c.expected)
		}
	}
}
