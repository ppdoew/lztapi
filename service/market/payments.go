package market

import (
	"LztApi/lztapi/dto"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func (m *Market) Payout() (*dto.Response, error) {
	return m.doRequest(resty.MethodGet, "balance/payout/services", nil, nil)
}

func (m *Market) CreatePayout(data any) (*dto.Response, error) {
	switch v := data.(type) {
	case map[string]interface{}:
	case dto.CreatePayoutOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип json: %T | Поддерживаемые: map[string]interface{}, dto.CreatePayoutOptional",
			v,
		)
	}

	return m.doRequest(resty.MethodPost, "balance/payout", nil, data)
}

func (m *Market) Transfer(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.TransferOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.TransferOptional",
			v,
		)
	}

	return m.doRequest(resty.MethodPost, "balance/transfer", params, nil)
}

func (m *Market) CancelTransfer(paymentId int) (*dto.Response, error) {
	return m.doRequest(resty.MethodPost, "balance/transfer/cancel", map[string]int{"paymentId": paymentId}, nil)
}

func (m *Market) PaymentHistory(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.PaymentHistoryOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.PaymentHistoryOptional", v)
	}
	return m.doRequest(resty.MethodGet, "user/payments", params, nil)
}

func (m *Market) CheckTransferFee(amount ...int) (*dto.Response, error) {
	var params map[string]string
	if len(amount) > 0 {
		params = map[string]string{"amount": strconv.Itoa(amount[0])}
	}

	return m.doRequest(resty.MethodGet, "balance/transfer/fee", params, nil)
}

func (m *Market) CurrencyList() (*dto.Response, error) {
	return m.doRequest(resty.MethodGet, "currency", nil, nil)
}

func (m *Market) AutoPayments() (*dto.Response, error) {
	return m.doRequest(resty.MethodGet, "auto-payments", nil, nil)
}

func (m *Market) CreateAutoPayments(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.CreateAutoPaymentsOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.CreateAutoPaymentsOptional", v)
	}
	return m.doRequest(resty.MethodPost, "auto-payment", params, nil)
}

func (m *Market) DeleteAutoPayments(autoPaymentsId int) (*dto.Response, error) {
	return m.doRequest(resty.MethodDelete, "auto-payments", map[string]int{"auto_payment_id": autoPaymentsId}, nil)
}
