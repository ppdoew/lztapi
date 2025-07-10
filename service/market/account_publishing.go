package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) FastSell(params any, body any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.FastSellOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.FastSellOptional", v)
	}
	switch v := body.(type) {
	case map[string]interface{}:
	case dto.FastSellBodyOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип body: %T | Поддерживаемые: map[string]interface{}, dto.FastSellBodyOptional", v)
	}
	return m.doRequest(resty.MethodPost, "item/fast-sell", params, body)
}

func (m *Market) AddAccount(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.AddAccountOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.AddAccountOptional",
			v,
		)
	}

	return m.doRequest(resty.MethodPost, "item/add", params, nil)
}

func (m *Market) CheckAccountUpload(
	itemID int,
	params any,
	body any,
) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.CheckAccountOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | поддерживаемые: map[string]string, dto.CheckAccountOptional",
			v,
		)
	}

	switch v := body.(type) {
	case map[string]interface{}:
	case dto.CheckAccountBodyOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип body: %T | поддерживаемые: map[string]interface{}, dto.CheckAccountBodyOptional",
			v,
		)
	}

	path := fmt.Sprintf("%d/goods/check", itemID)
	return m.doRequest(resty.MethodPost, path, params, body)
}
