# Go Enswitch

The unofficial [Enswitch](https://integrics.com/enswitch/) Go client library.

## Requirements
- Go 1.15 or later

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):

``` shell
go mod init
```

Then, reference enswitch-go in a Go program with `import`:

``` go
import (
	"github.com/aleyrizvi/enswitch-go/v1"
	"github.com/aleyrizvi/enswitch-go/v1/customer"
)
```

## Future statuses

![50%](https://progress-bar.dev/10?title=Progress)

| Status | Module name   | Description                      |
|--------|---------------|----------------------------------|
|        | access        | Remote access                    |
|        | addresses     | SIP addresses                    |
|        | alerts        | Alerts                           |
|        | audit         | Audit log                        |
|        | bulk          | Bul dialer                       |
|        | calls         | Calls                            |
|        | callshops     | Callshops                        |
|        | cards         | Calling cards                    |
|        | cdrs          | Call history                     |
|        | charges       | Charges                          |
|        | classes       | Number classes                   |
|        | codes         | Feature codes                    |
|        | conferences   | Conferences                      |
|        | configs       | Configuration settings           |
|        | cos           | Classes of service               |
|        | customers     | Customers                        |
|        | customs       | Custom settings                  |
|        | domains       | Domains                          |
|        | emails        | Emails                           |
|        | events        | Events                           |
|        | faxes         | Faxes                            |
|        | features      | Telephony features               |
|        | hotels        | Hotel gateways                   |
|        | huntgroups    | Hunt groups                      |
|        | incosts       | Costs for inbound calls          |
|        | ingroups      | Inbound groups                   |
|        | invoices      | Invoices                         |
|        | ivrs          | IVR menus                        |
|        | machines      | Machines                         |
|        | mailboxes     | Mailboxes                        |
|        | menus         | Menus                            |
|        | messages      | Messages                         |
|        | music         | Music                            |
|        | notes         | Notes                            |
|        | numbers       | Numbers                          |
|        | nvgroups      | Number vendor groups             |
|        | outcosts      | Costs for outbound calls         |
|        | outgroups     | Outbound groups                  |
|        | pagegroups    | Page groups                      |
|        | patternmenus  | Pattern menus                    |
|        | peers         | Peers                            |
|        | people        | People                           |
|        | phones        | Telephone lines                  |
|        | pickupgroups  | Pickup groups                    |
|        | plans         | Rate plans                       |
|        | plugins       | Plugins                          |
|        | preferences   | Preferences                      |
|        | prefixes      | Feature prefixes                 |
|        | products      | Products                         |
|        | provisioning  | Provisioning templates and files |
|        | queues        | Queues                           |
|        | recording     | Call recording                   |
|        | regions       | Regions                          |
|        | requests      | Requests                         |
|        | roles         | Roles                            |
|        | routes        | Routes                           |
|        | sounds        | Sounds                           |
|        | speeddials    | Speed dials                      |
|        | subscriptions | Subscriptions                    |
|        | taxes         | Taxes                            |
|        | timegroups    | Time groups                      |
|        | traces        | Traces                           |
|        | transactions  | Transactions                     |
|        | unlimited     | Unlimited access                 |
|        | user          | Web user                         |
|        | vouchers      | Vouchers                         |

# Supported providers 
#### (Send a pull request to add your name)
- [Nomado telecom SPRL](https://www.nomado.eu)

# Support
I run on coffee. If you like to support me move faster, please help :)

<a href="https://www.buymeacoffee.com/asadrizvi" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
