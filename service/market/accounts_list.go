package market

import (
	"LztApi/lztapi/dto"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (m *Market) AllUserAccounts(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.AllUserAccountsOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый AllUserAccountsOptional", v)
	}
	return m.doRequest(resty.MethodGet, "user/items", params, nil)
}

func (m *Market) AllUserPurchased(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.AllUserPurchasedOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый AllUserPurchasedOptional", v)
	}
	return m.doRequest(resty.MethodGet, "user/orders", params, nil)
}

func (m *Market) AllFavouritesAccounts(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.AllFavouritesAccountsOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый AllFavouritesAccountsOptional", v)
	}
	return m.doRequest(resty.MethodGet, "fave", params, nil)
}

func (m *Market) AllViewedAccounts(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.AllViewedAccountsOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип params: %T | Поддерживаемый AllViewedAccountsOptional", v)
	}
	return m.doRequest(resty.MethodGet, "viewed", params, nil)
}
