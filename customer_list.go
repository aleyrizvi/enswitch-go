package enswitch

import (
	"context"
	"fmt"
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

func (c *Customer) List(ctx context.Context, input *CustomerListParams) {
	req, err := c.client.newRequest(ctx, "customers/list", getCustomerListOptions(input))
	if err != nil {
		return
	}

	fmt.Println(req.URL.String())
}

func getCustomerListOptions(input *CustomerListParams) url.Values {
	var qs url.Values

	return qs
}
