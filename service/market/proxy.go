package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) Proxy() (*dto.Response, error) {
	return m.doRequest(resty.MethodGet, "proxy", nil, nil)
}

func (m *Market) AddProxy(value interface{}) (*dto.Response, error) {
	switch v := value.(type) {
	case map[string]interface{}:
	case dto.AddProxyOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип json: %T | Поддерживаемый AddProxyOptional", v)
	}
	return m.doRequest(resty.MethodPost, "proxy", value, nil)
}

func (m *Market) DeleteProxy(value interface{}) (*dto.Response, error) {
	switch v := value.(type) {
	case map[string]interface{}:
	case dto.DeleteProxyOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип json: %T | Поддерживаемый DeleteProxyProxy", v)
	}
	return m.doRequest(resty.MethodDelete, "proxy", value, nil)
}
