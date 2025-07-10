package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) Profile() (*dto.Response, error) {
	return m.doRequest(resty.MethodGet, "me", nil, nil)
}

func (m *Market) EditMarketSetting(value any) (*dto.Response, error) {
	switch v := value.(type) {
	case map[string]interface{}:
	case dto.EditMarketSettingsOptional:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип json: %T | Поддерживаемый EditMarketSettingsOptional", v)
	}
	return m.doRequest(resty.MethodPut, "me", value, nil)
}
