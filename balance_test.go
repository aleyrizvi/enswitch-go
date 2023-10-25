package enswitch

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

		if err != c.errorExpected {
			t.Errorf("[balanceUpdate]: given %v, got %v", c.errorExpected, err)
		}

		if res.Responses[0].Code != 204 {
			t.Error("[balanceUpdate]: status code does not match")
		}
	}
}
