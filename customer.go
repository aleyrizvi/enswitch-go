package enswitch

type (
	billType               string
	customerListSort       string
	balanceTransactionType string
)

const (
	BillTypeCallShop        billType = "callshop"
	BillTypeExternal        billType = "external"
	BillTypeNone            billType = "none"
	BillTypePostPaid        billType = "postpaid"
	BillTypePrepaid         billType = "prepaid"
	BillTypePrepaidCalls    billType = "prepaid_calls"
	BillTypePrepaidInvoices billType = "prepaid_invoices"

	CustomerListSortByAccount     customerListSort = "account"
	CustomerListSortByDescription customerListSort = "description"
	CustomerListSortByName        customerListSort = "name"

	TransactionTypeCredit balanceTransactionType = "credit"
	TransactionTypeTopUp  balanceTransactionType = "topup"
)

type Customer struct {
	client  *Client
	Balance *Balance
}
