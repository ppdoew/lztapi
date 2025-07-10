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

### Customizing Client Behavior

```go
// Change rate limiter
client.SetLimiter(rate.NewLimiter(rate.Every(100*time.Millisecond), 1))

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
    {Uri: "https://prod-api.lzt.market/user/payments", Method: dto.MethodGet},
    {Uri: "https://prod-api.lzt.market/letters2", Method: dto.MethodGet, Params: dto.GetLettersParams{ItemID: 123, Limit: 20}},
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

## Contributing

Contributions are welcome! Please open issues and submit pull requests.

---

*Generated from [ppdoew/lztapi](https://github.com/ppdoew/lztapi)*  
