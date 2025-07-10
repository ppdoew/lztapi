# Клиент LZT API для Go

Библиотека клиента на Go для взаимодействия с API **LZT Market**. Предоставляет поддержку:
- Аутентификация с помощью JWT-токенов
- Ограничение частоты запросов
- Логика повторных попыток
- Настройка прокси
- Статически типизированные DTO в пакете `dto`
- Пакетные (batch) запросы
- Единая структура ответа

## Установка

```bash
go get github.com/ppdoew/lztapi
```

## Начало работы

### Создание сессии клиента

```go
import (
  "github.com/ppdoew/lztapi"
)

client := lztapi.CreateClient("YOUR_JWT_TOKEN")
```

Это инициализирует:

- HTTP-клиент Resty
- Заголовок `Content-Type: application/json`
- Заголовок `Authorization: Bearer <token>`
- Ограничитель скорости (по умолчанию: 1 запрос каждые 200 мс)
- Политику повторных попыток (конфигурируется через `config` по умолчанию)
- Таймаут (конфигурируется через `config.DefaultTimeout`)

### Конфигурация прокси

Вы можете направлять запросы через прокси:

```go
// Поддерживается HTTP, HTTPS и SOCKS5.
err := client.SetProxy("http://127.0.0.1:8080")
if err != nil {
  // обработать ошибку
}
```

С логином и паролем:

```go
// Поддерживается HTTP, HTTPS и SOCKS5.
err := client.SetProxy("http://login:password@127.0.0.1:8080")
if err != nil {
  // обработать ошибку
}
```

### Настройка поведения клиента

```go
// Изменение ограничителя скорости
// 1 запрос каждые 200 миллисекунд
client.SetLimiter(rate.NewLimiter(rate.Every(200*time.Millisecond), 1))

// Настройка таймаутов и политики повторных попыток
client.SetTimeout(10 * time.Second)
client.SetRetryCount(5)
client.SetRetryWaitTime(1 * time.Second)
client.SetRetryMaxWaitTime(5 * time.Second)
```

## DTOs

Все типы запросов и ответов API определены в пакете `dto`.  
Ключевые типы включают:
- `dto.MethodGet`, `dto.MethodPost` и т.д.
- `dto.BatchJob` для пакетных запросов
- Специфичные для модуля DTO, например `dto.Payment`, `dto.Letter` и т.д.

## Модули

Библиотека клиента организована по модулям. Например, модуль Market:

```go
import (
  "github.com/ppdoew/lztapi"
)

apiClient := lztapi.CreateClient("TOKEN")
marketClient := lztapi.ModuleMarket(apiClient)
```

## Пакетные запросы

Выполняйте до 10 запросов API в одном HTTP-запросе:

```go
jobs := []dto.BatchJob{
    {Uri: "https://prod-api.lzt.market/user/payments", Method: dto.MethodGet, Params: dto.PaymentHistoryOptional{Type: dto.OpPaidItem}},
    {Uri: "https://prod-api.lzt.market/letters2", Method: dto.MethodGet},
}

batchResp, err := marketClient.Batch(jobs)
if err != nil {
  // обработать ошибку
}

// Просмотр результатов отдельных запросов:
for _, job := range batchResp.Jobs {
    fmt.Println("Status:", job.Response.StatusCode)
    fmt.Println("Body:", string(job.Response.Body))
}
```

## Единая структура ответа

Все ответы используют общий `Response`:

```go
type Response struct {
  StatusCode int                    `json:"status_code"`
  Body       []byte                 `json:"body"`
  Text       string                 `json:"text"`
  Json       map[string]interface{} `json:"json"`
}
```

- `StatusCode`: HTTP-код ответа
- `Body`: Сырые байты тела ответа
- `Text`: Тело в виде строки
- `Json`: Распарсенный JSON-объект

## Как использовать клиент

Структура клиента:

```go
type LztClient struct {
  // HTTP-клиент с аутентификацией, таймаутами и логированием
  Client  *resty.Client

  // Ограничитель скорости для контроля частоты запросов
  Limiter *rate.Limiter

  // Базовый URL API (например, https://prod-api.lzt.market/)
  baseURL string
}
```

Пример:
```go
client := lztapi.CreateClient("TOKEN")
resp, err := client.Client.R().Get("me")
if err != nil {

}

fmt.Println(resp.String())
```

## Методы модуля Market

### Управление аккаунтом
- `GetAccount(itemID int, ParseSameItemIds ...bool) (*dto.Response, error)`
- `BulkGetAccounts(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.BulkGetAccountsOptional`
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

### Управление Telegram
- `TelegramConfirmationCode(itemID int) (*dto.Response, error)`
- `TelegramResetAuth(itemID int) (*dto.Response, error)`

### Интеграция с Steam
- `AccountSteamInventoryValue(itemID int, params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.SteamInventoryValueOptional`
- `SteamInventoryValue(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.SteamInventoryValueOptional`
- `SteamHTML(itemID int, htmlType dto.TypeHtml) (*dto.Response, error)`
- `UpdateInventoryValue(itemID int, params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.SteamInventoryValueOptional`
- `GetMaFile(itemID int) (*dto.Response, error)`
- `RemoveMaFile(itemID int) (*dto.Response, error)`
- `ConfirmSDA(itemID int, params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.ConfirmSDAOptional`
- `MaFileConfirmationCode(itemID int) (*dto.Response, error)`

### Списки аккаунтов пользователей
- `AllUserAccounts(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.AllUserAccountsOptional`
- `AllUserPurchased(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.AllUserPurchasedOptional`
- `AllFavouritesAccounts(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.AllFavouritesAccountsOptional`
- `AllViewedAccounts(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.AllViewedAccountsOptional`

### Публикация аккаунтов
- `FastSell(params any, body any) (*dto.Response, error)`
  - Params: `map[string]string`, `dto.FastSellOptional`  
  - Body: `map[string]interface{}`, `dto.FastSellBodyOptional`
- `AddAccount(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.AddAccountOptional`
- `CheckAccountUpload(itemID int, params any, body any) (*dto.Response, error)`
  - Params: `map[string]string`, `dto.CheckAccountUploadOptional`  
  - Body: `map[string]interface{}`, `dto.CheckAccountUploadBodyOptional`

### Покупка аккаунтов
- `FastBuy(itemID int, price ...int) (*dto.Response, error)`
- `CheckAccount(itemID int) (*dto.Response, error)`
- `ConfirmBuy(itemID int, price ...int) (*dto.Response, error)`

### Категории и поиск
- `CategorySearch(category dto.Category, params map[string]string) (*dto.Response, error)`
- `CategoryGames(category dto.Category) (*dto.Response, error)`
- `CategorySearchParams(category dto.Category) (*dto.Response, error)`
- `Categories(topQueries ...bool) (*dto.Response, error)`

### Платежи
- `Payout() (*dto.Response, error)`
- `CreatePayout(data any) (*dto.Response, error)`
  - Принимает: `map[string]interface{}`, `dto.CreatePayoutOptional`
- `Transfer(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.TransferOptional`
- `CancelTransfer(paymentId int) (*dto.Response, error)`
- `PaymentHistory(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.PaymentHistoryOptional`
- `CheckTransferFee(amount ...int) (*dto.Response, error)`
- `CurrencyList() (*dto.Response, error)`
- `AutoPayments() (*dto.Response, error)`
- `CreateAutoPayments(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.CreateAutoPaymentsOptional`
- `DeleteAutoPayments(autoPaymentsId int) (*dto.Response, error)`

### Профиль
- `Profile() (*dto.Response, error)`
- `EditMarketSetting(value any) (*dto.Response, error)`
  - Принимает: `map[string]interface{}`, `dto.EditMarketSettingsOptional`

### Прокси
- `Proxy() (*dto.Response, error)`
- `AddProxy(value interface{}) (*dto.Response, error)`
  - Принимает: `map[string]interface{}`, `dto.AddProxyOptional`
- `DeleteProxy(value interface{}) (*dto.Response, error)`
  - Принимает: `map[string]interface{}`, `dto.DeleteProxyOptional`

### Счета-фактуры
- `InvoiceList(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.InvoiceListOptional`
- `Invoice(params any) (*dto.Response, error)`
  - Принимает: `map[string]string`, `dto.InvoiceOptional`
- `CreateInvoice(value any) (*dto.Response, error)`
  - Принимает: `map[string]interface{}`, `dto.CreateInvoiceOptional`
