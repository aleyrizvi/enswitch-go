package enswitch

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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

func TestCustomerBalanceQueryParams(t *testing.T) {
	type Test struct {
		b        BalanceParams
		expected url.Values
	}

	transaction := TransactionTypeTopUp

	cases := []Test{
		{
			b: BalanceParams{
				ID:     1,
				Amount: 20,
			},
			expected: map[string][]string{
				"id":     {"1"},
				"amount": {"20"},
			},
		},
		{
			b: BalanceParams{
				ID:          1,
				Amount:      20,
				Currency:    String("eur"),
				Transaction: &transaction,
				Description: String("test"),
			},
			expected: map[string][]string{
				"id":          {"1"},
				"amount":      {"20"},
				"currency":    {"eur"},
				"transaction": {"topup"},
				"description": {"test"},
			},
		},
	}

	for _, c := range cases {
		qs := balanceQueryParam(c.b)

		if qs.Encode() != c.expected.Encode() {
			t.Errorf("[Balance] unable to create query param. Given %v, got %v", qs, c.expected)
		}

	}
}

func TestBalance_Update1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx   context.Context
		input BalanceParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BalanceResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Balance{
				client: tt.fields.client,
			}
			got, err := b.Update(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_Update2(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx   context.Context
		input BalanceParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BalanceResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Balance{
				client: tt.fields.client,
			}
			got, err := b.Update(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_call(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		username   string
		password   string
		baseURL    string
		Customer   *Customers
	}
	type args struct {
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				username:   tt.fields.username,
				password:   tt.fields.password,
				baseURL:    tt.fields.baseURL,
				Customer:   tt.fields.Customer,
			}
			got, err := c.call(tt.args.req, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("call() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_newRequest(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		username   string
		password   string
		baseURL    string
		Customer   *Customers
	}
	type args struct {
		ctx context.Context
		uri string
		qs  url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				username:   tt.fields.username,
				password:   tt.fields.password,
				baseURL:    tt.fields.baseURL,
				Customer:   tt.fields.Customer,
			}
			got, err := c.newRequest(tt.args.ctx, tt.args.uri, tt.args.qs)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_parseQueryParams(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		username   string
		password   string
		baseURL    string
		Customer   *Customers
	}
	type args struct {
		qs url.Values
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				username:   tt.fields.username,
				password:   tt.fields.password,
				baseURL:    tt.fields.baseURL,
				Customer:   tt.fields.Customer,
			}
			if got := c.parseQueryParams(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseQueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomers_Get(t *testing.T) {
	type fields struct {
		client  *Client
		Balance *Balance
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customers{
				client:  tt.fields.client,
				Balance: tt.fields.Balance,
			}
			c.Get()
		})
	}
}

func TestCustomers_List(t *testing.T) {
	type fields struct {
		client  *Client
		Balance *Balance
	}
	type args struct {
		ctx   context.Context
		input *CustomerListParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customers{
				client:  tt.fields.client,
				Balance: tt.fields.Balance,
			}
			c.List(tt.args.ctx, tt.args.input)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		username string
		password string
		baseURL  string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.username, tt.args.password, tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_balanceQueryParam(t *testing.T) {
	type args struct {
		b BalanceParams
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceQueryParam(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCustomerListOptions(t *testing.T) {
	type args struct {
		in0 *CustomerListParams
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCustomerListOptions(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCustomerListOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
