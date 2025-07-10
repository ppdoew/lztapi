package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
	"strconv"
)

func (m *Market) FastBuy(itemId int, price ...int) (*dto.Response, error) {
	var queryParams map[string]string
	if len(price) > 0 {
		queryParams = map[string]string{
			"price": strconv.Itoa(price[0]),
		}
	}

	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/fast-buy", itemId),
		queryParams,
		nil,
	)
}

func (m *Market) CheckAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/check-account", itemId),
		nil,
		nil,
	)
}

func (m *Market) ConfirmBuy(itemId int, price ...int) (*dto.Response, error) {
	var queryParams map[string]string
	if len(price) > 0 {
		queryParams = map[string]string{
			"price": strconv.Itoa(price[0]),
		}
	}

	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/confirm-buy", itemId),
		queryParams,
		nil,
	)
}
