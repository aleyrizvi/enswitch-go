package enswitch

import (
	"context"
	"net/url"
)

type CustomerListParams struct {
	Deleted    *bool
	BillType   *billType
	Count      *bool
	Descending *bool
	Directory  *uint32
	Limit      *uint32
	Name       *string
	Offset     *uint32
	Pages      *uint32
	Parent     *uint32
	Recursive  *bool
}

func (c *Customers) List(ctx context.Context, input *CustomerListParams) {
	_, err := c.client.newRequest(ctx, "customers/list", getCustomerListOptions(input))
	if err != nil {
		return
	}
}

func getCustomerListOptions(_ *CustomerListParams) url.Values {
	var qs url.Values

	return qs
}
