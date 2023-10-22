package enswitch

import (
	"context"
	"fmt"
	"net/url"
)

type BalanceParams struct {
	ID          uint32
	Amount      float64
	Currency    *string
	Transaction *balanceTransactionType
	Description *string
}

type BalanceResponse Response

type Balance struct {
	client *Client
}

func (b *Balance) Update(ctx context.Context, input BalanceParams) (*BalanceResponse, error) {
	qs := customerBalanceQueryParams(input)

	req, err := b.client.newRequest(ctx, "customers/balance/add", qs)
	if err != nil {
		return nil, err
	}

	var response *BalanceResponse

	res, err := b.client.call(req, &response)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return response, nil
}

func customerBalanceQueryParams(b BalanceParams) url.Values {
	qs := url.Values{}

	qs.Add("id", fmt.Sprintf("%v", b.ID))
	qs.Add("amount", fmt.Sprintf("%v", b.Amount))

	if b.Currency != nil {
		qs.Add("currency", *b.Currency)
	}

	if b.Transaction != nil {
		qs.Add("transaction", string(*b.Transaction))
	}

	if b.Description != nil {
		qs.Add("description", *b.Description)
	}

	return qs
}
