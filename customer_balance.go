package enswitch

import (
	"context"
	"fmt"
	"net/url"
)

type CustomerBalanceParams struct {
	ID          uint32
	Amount      float64
	Currency    *string
	Transaction *balanceTransactionType
	Description *string
}

func (b *Customer) UpdateBalance(ctx context.Context, input CustomerBalanceParams) (*Response, error) {
	qs := customerBalanceQueryParams(input)

	req, err := b.client.newRequest(ctx, "customers/balance/add", qs)

	if err != nil {
		return nil, err
	}

	var response *Response

	_, err = b.client.call(req, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func customerBalanceQueryParams(b CustomerBalanceParams) url.Values {
	qs := url.Values{}

	qs.Add("id", fmt.Sprintf("%v", b.ID))
	qs.Add("amount", fmt.Sprintf("%v", b.Amount))

	if b.Currency != nil {
		qs.Add("currency", *b.Currency)
	}

	if b.Transaction != nil {
		qs.Add("transaction", fmt.Sprintf("%s", *b.Transaction))
	}

	if b.Description != nil {
		qs.Add("description", fmt.Sprintf("%s", *b.Description))
	}

	return qs
}
