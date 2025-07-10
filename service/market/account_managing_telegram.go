package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) TelegramConfirmationCode(itemID int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/telegram-login-code", itemID),
		nil,
		nil,
	)
}

func (m *Market) TelegramResetAuth(itemID int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/telegram-reset-authorizations", itemID),
		nil,
		nil,
	)
}
