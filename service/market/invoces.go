package market

import (
	"LztApi/lztapi/dto"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (m *Market) InvoiceList(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.InvoiceListOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый InvoiceListOptional", v)
	}
	return m.doRequest(resty.MethodGet, "invoice/list", params, nil)
}

func (m *Market) Invoice(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.InvoiceOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый InvoiceOptional", v)
	}
	return m.doRequest(resty.MethodGet, "invoice", params, nil)
}

func (m *Market) CreateInvoice(value any) (*dto.Response, error) {
	switch v := value.(type) {
	case map[string]interface{}:
	case dto.CreateInvoiceOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип json: %T | Поддерживаемый CreateInvoiceOptional", v)
	}
	return m.doRequest(resty.MethodPost, "invoice", value, nil)
}
