# LZT API Go Client

This Go module provides a client for interacting with the LOLZTEAM Public API (LZT Market). It supports rate limiting, retries, proxy configuration, DTOs, batch requests, and unified response handling.

## Installation

```bash
go get github.com/yourusername/lztapi
```

## Usage

```go
import (
    "LztApi/lztapi"
    "LztApi/lztapi/dto"
)

func main() {
    // Create client with JWT token
    client := lztapi.CreateClient("YOUR_JWT_TOKEN")

    // Configure proxy (optional)
    err := client.SetProxy("http://127.0.0.1:8080")
    if err != nil {
        panic(err)
    }

    // Change rate limiter, timeouts, retries as needed
    client.SetTimeout(10 * time.Second)
    client.SetRetryCount(3)

    // Initialize module
    market := lztapi.ModuleMarket(client)

    // DTOs for batch jobs
    jobs := []dto.BatchJob{
        {URI: "https://prod-api.lzt.market/user/payments", Method: dto.MethodGet},
    }

    // Perform batch request
    resp, err := market.Batch(jobs)
    if err != nil {
        panic(err)
    }

    // Inspect unified response
    for _, job := range resp.Jobs {
        fmt.Printf("Job ID: %s, Status: %d, Body: %s\n", job.JobID, job.Response.StatusCode, job.Response.Text)
    }
}
```

## Client Creation

The client is created using a JWT token:

```go
client := lztapi.CreateClient("YOUR_JWT_TOKEN")
```

It configures:
- **Rate Limiter**: 5 requests per second by default.
- **Timeout**: default from `config.DefaultTimeout`.
- **Retries**: count and backoff from config.

### Proxy Support

You can set an HTTP, HTTPS, or SOCKS5 proxy:

```go
err := client.SetProxy("http://proxy.example.com:3128")
```

Pass an empty string to disable proxy:

```go
client.SetProxy("")
```

Unsupported schemes return an error.

## DTOs

All request and response types are defined in the `dto` package. Key types include:

- `dto.BatchJob`:
    ```go
    type BatchJob struct {
        Id     string        `json:"id,omitempty"`
        Uri    string        `json:"uri"`
        Method MethodRequest `json:"method,omitempty"`
        Params interface{}   `json:"params,omitempty"`
    }
    ```
- HTTP methods:
  ```go
  type HTTPMethod string
  const (
      MethodGet    HTTPMethod = "GET"
      MethodPost   HTTPMethod = "POST"
      MethodPut    HTTPMethod = "PUT"
      MethodDelete HTTPMethod = "DELETE"
  )
  ```

## Batch Requests

Use the `/batch` endpoint to perform up to 10 requests in one call:

```go
jobs := []dto.BatchJob{
    {URI: "https://prod-api.lzt.market/user/payments", Method: dto.MethodGet},
    {URI: "https://prod-api.lzt.market/letters2", Method: dto.MethodGet, Params: dto.GetLettersParams{ItemID: 123, Limit: 20}},
}

resp, err := market.Batch(jobs)
```

## Unified Response Schema

All methods return a unified `Response` struct:

```go
type Response struct {
    StatusCode int                    `json:"status_code"`
    Body       []byte                 `json:"body"`
    Text       string                 `json:"text"`
    JSON       map[string]interface{} `json:"json"`
}
```

- **StatusCode**: HTTP status code.
- **Body**: Raw response bytes.
- **Text**: Raw response body as string.
- **JSON**: Parsed JSON into a generic map.

This allows consistent handling of all API calls.

## Configuration

Defaults are set in `lztapi/config`:

- `DefaultBaseURL` – base API URL.
- `DefaultTimeout` – HTTP client timeout.
- `DefaultRetryCount` – number of retry attempts.
- `DefaultRetryWait` – initial retry wait time.
- `DefaultRetryMaxWait` – maximum retry wait time.
