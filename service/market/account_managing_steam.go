package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) AccountSteamInventoryValue(itemID int, params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string, dto.SteamInventoryValueOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.SteamInventoryValueOptional",
			v,
		)
	}

	path := fmt.Sprintf("%d/inventory-value", itemID)
	return m.doRequest(
		resty.MethodGet,
		path,
		params,
		nil,
	)
}

func (m *Market) SteamInventoryValue(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.SteamValueOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.SteamValueOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodGet,
		"steam-value",
		params,
		nil,
	)
}

func (m *Market) SteamHTML(itemID int, htmlType dto.TypeHtml) (*dto.Response, error) {
	switch htmlType {
	case dto.TypeProfiles,
		dto.TypeGames:
	default:
		return nil, fmt.Errorf(
			"неверный htmlType %q: должен быть %q или %q",
			htmlType, dto.TypeProfiles, dto.TypeGames,
		)
	}

	params := map[string]string{}
	if string(htmlType) != "" {
		params["type"] = string(htmlType)
	}

	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/steam-preview", itemID),
		params,
		nil,
	)
}

func (m *Market) UpdateInventoryValue(itemID int, params any) (*dto.Response, error) {
	switch v := params.(type) {
	case map[string]string:
	case dto.UpdateSteamValueOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.SteamUpdateValueOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/update-inventory", itemID),
		params,
		nil,
	)
}

func (m *Market) GetMaFile(itemID int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/mafile", itemID),
		nil,
		nil,
	)
}

func (m *Market) RemoveMaFile(itemID int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodDelete,
		fmt.Sprintf("%d/mafile", itemID),
		nil,
		nil,
	)
}

func (m *Market) ConfirmSDA(itemID int, params any) (*dto.Response, error) {
	switch v := params.(type) {
	case nil:
	case map[string]string:
	case dto.ConfirmSDAOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.ConfirmSDAOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/confirm-sda", itemID),
		params,
		nil,
	)
}

func (m *Market) MaFileConfirmationCode(itemID int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/guard-code", itemID),
		nil,
		nil,
	)
}
