# LZT API Go Client

A Go client library for interacting with the **LZT Market** API. Provides support for:
- Authentication with JWT tokens
- Rate limiting
- Retry logic
- Proxy configuration
- Strongly-typed DTOs in the `dto` package
- Batch requests
- Unified response schema

## Installation

```bash
go get github.com/ppdoew/lztapi
```

## Getting Started

### Creating a Client Session

```go
import (
  "github.com/ppdoew/lztapi"
)

client := lztapi.CreateClient("YOUR_JWT_TOKEN")
```

This will initialize:

- A Resty HTTP client
- JSON `Content-Type` header
- `Authorization: Bearer <token>`
- Rate limiter (default: 1 request per 200ms)
- Retry policy (configured via `config` defaults)
- Timeout (configured via `config.DefaultTimeout`)

### Proxy Configuration

You can route requests through a proxy:

```go
// Accepts HTTP, HTTPS or SOCKS5 proxies.
err := client.SetProxy("http://127.0.0.1:8080")
if err != nil {
  // handle error
}
```

With login:password

```go
// Accepts HTTP, HTTPS or SOCKS5 proxies.
err := client.SetProxy("http://login:password@127.0.0.1:8080")
if err != nil {
  // handle error
}
```

### Customizing Client Behavior

```go
// Change rate limiter
// 1 request every 200 milliseconds
client.SetLimiter(rate.NewLimiter(rate.Every(200*time.Millisecond), 1))

// Adjust timeouts and retry policy
client.SetTimeout(10 * time.Second)
client.SetRetryCount(5)
client.SetRetryWaitTime(1 * time.Second)
client.SetRetryMaxWaitTime(5 * time.Second)
```

## DTOs

All API request and response types are defined in the `dto` package.  
Key types include:
- `dto.MethodGet`, `dto.MethodPost`, etc.
- `dto.BatchJob` for batch requests
- Module-specific DTOs, e.g. `dto.Payment`, `dto.Letter`, etc.

## Modules

The client library is organized into modules. For example, the Market module:

```go
import (
  "github.com/ppdoew/lztapi"
)

apiClient := lztapi.CreateClient("TOKEN")
marketClient := lztapi.ModuleMarket(apiClient)
```

## Batch Requests

Execute up to 10 API calls in a single HTTP request:

```go
jobs := []dto.BatchJob{
    {Uri: "https://prod-api.lzt.market/user/payments", Method: dto.MethodGet, Params: dto.PaymentHistoryOptional{Type: dto.OpPaidItem}},
    {Uri: "https://prod-api.lzt.market/letters2", Method: dto.MethodGet},
}

batchResp, err := marketClient.Batch(jobs)
if err != nil {
  // handle error
}

// Inspect individual job results:
for _, job := range batchResp.Jobs {
    fmt.Println("Status:", job.Response.StatusCode)
    fmt.Println("Body:", string(job.Response.Body))
}
```

## Unified Response Schema

All responses use a common `Response` struct:

```go
type Response struct {
  StatusCode int                    `json:"status_code"`
  Body       []byte                 `json:"body"`
  Text       string                 `json:"text"`
  Json       map[string]interface{} `json:"json"`
}
```

- `StatusCode`: HTTP status code
- `Body`: Raw bytes of the response body
- `Text`: Body as string
- `Json`: Unmarshaled JSON object

## How to Use the Client

Struct client:

```go
type LztClient struct {
  // HTTP client with authentication, timeouts, and logging
  Client  *resty.Client
  
  // Rate limiter for controlling request throughput
  Limiter *rate.Limiter
  
  // Base API URL (for example, https://prod-api.lzt.market/)
  baseURL string
}
```
Example:
```go
client := lztapi.CreateClient("TOKEN")
resp, err := client.Client.R().Get("me")
if err != nil {
	
}

fmt.Println(resp.String())
```

## Market Module Methods

### Account Management
- `GetAccount(itemID int, ParseSameItemIds ...bool) (*dto.Response, error)`
- `BulkGetAccounts(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.BulkGetAccountsOptional`
- `EditAccount(itemID int, query any, body any) (*dto.Response, error)`
  - Query: `map[string]string`, `dto.EditAccountQueryOptional`  
  - Body: `map[string]interface{}`, `dto.EditAccountBodyOptional`
- `DeleteAccount(itemID int, reason string) (*dto.Response, error)`
- `BumpAccount(itemID int) (*dto.Response, error)`
- `OpenAccount(itemID int) (*dto.Response, error)`
- `CloseAccount(itemID int) (*dto.Response, error)`
- `CreateArbitrage(itemID int, postBody string) (*dto.Response, error)`
- `EditNote(itemID int, text string) (*dto.Response, error)`
- `AccountImage(itemID int, imageType string) (*dto.Response, error)`

### Telegram Management
- `TelegramConfirmationCode(itemID int) (*dto.Response, error)`
- `TelegramResetAuth(itemID int) (*dto.Response, error)`

### Steam Integration
- `AccountSteamInventoryValue(itemID int, params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.SteamInventoryValueOptional`
- `SteamInventoryValue(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.SteamInventoryValueOptional`
- `SteamHTML(itemID int, htmlType dto.TypeHtml) (*dto.Response, error)`
- `UpdateInventoryValue(itemID int, params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.SteamInventoryValueOptional`
- `GetMaFile(itemID int) (*dto.Response, error)`
- `RemoveMaFile(itemID int) (*dto.Response, error)`
- `ConfirmSDA(itemID int, params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.ConfirmSDAOptional`
- `MaFileConfirmationCode(itemID int) (*dto.Response, error)`

### User Accounts Lists
- `AllUserAccounts(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.AllUserAccountsOptional`
- `AllUserPurchased(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.AllUserPurchasedOptional`
- `AllFavouritesAccounts(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.AllFavouritesAccountsOptional`
- `AllViewedAccounts(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.AllViewedAccountsOptional`

### Account Publishing
- `FastSell(params any, body any) (*dto.Response, error)`
  - Params: `map[string]string`, `dto.FastSellOptional`  
  - Body: `map[string]interface{}`, `dto.FastSellBodyOptional`
- `AddAccount(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.AddAccountOptional`
- `CheckAccountUpload(itemID int, params any, body any) (*dto.Response, error)`
  - Params: `map[string]string`, `dto.CheckAccountUploadOptional`  
  - Body: `map[string]interface{}`, `dto.CheckAccountUploadBodyOptional`

### Account Purchasing
- `FastBuy(itemID int, price ...int) (*dto.Response, error)`
- `CheckAccount(itemID int) (*dto.Response, error)`
- `ConfirmBuy(itemID int, price ...int) (*dto.Response, error)`

### Categories and Search
- `CategorySearch(category dto.Category, params map[string]string) (*dto.Response, error)`
- `CategoryGames(category dto.Category) (*dto.Response, error)`
- `CategorySearchParams(category dto.Category) (*dto.Response, error)`
- `Categories(topQueries ...bool) (*dto.Response, error)`

### Payments Methods
- `Payout() (*dto.Response, error)`
- `CreatePayout(data any) (*dto.Response, error)`
  - Accepts: `map[string]interface{}`, `dto.CreatePayoutOptional`
- `Transfer(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.TransferOptional`
- `CancelTransfer(paymentId int) (*dto.Response, error)`
- `PaymentHistory(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.PaymentHistoryOptional`
- `CheckTransferFee(amount ...int) (*dto.Response, error)`
- `CurrencyList() (*dto.Response, error)`
- `AutoPayments() (*dto.Response, error)`
- `CreateAutoPayments(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.CreateAutoPaymentsOptional`
- `DeleteAutoPayments(autoPaymentsId int) (*dto.Response, error)`

### Profile
- `Profile() (*dto.Response, error)`
- `EditMarketSetting(value any) (*dto.Response, error)`
  - Accepts: `map[string]interface{}`, `dto.EditMarketSettingsOptional`

### Proxy
- `Proxy() (*dto.Response, error)`
- `AddProxy(value interface{}) (*dto.Response, error)`
  - Accepts: `map[string]interface{}`, `dto.AddProxyOptional`
- `DeleteProxy(value interface{}) (*dto.Response, error)`
  - Accepts: `map[string]interface{}`, `dto.DeleteProxyOptional`

### Invoices
- `InvoiceList(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.InvoiceListOptional`
- `Invoice(params any) (*dto.Response, error)`
  - Accepts: `map[string]string`, `dto.InvoiceOptional`
- `CreateInvoice(value any) (*dto.Response, error)`
  - Accepts: `map[string]interface{}`, `dto.CreateInvoiceOptional`
