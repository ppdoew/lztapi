package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
	"strconv"
)

func (m *Market) GetAccount(itemID int, ParseSameItemIds ...bool) (*dto.Response, error) {
	var queryParams map[string]string
	if len(ParseSameItemIds) > 0 {
		queryParams = map[string]string{
			"price": strconv.FormatBool(ParseSameItemIds[0]),
		}
	}

	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d", itemID),
		queryParams,
		nil,
	)
}

func (m *Market) BulkGetAccounts(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case nil:
	case map[string]interface{}:
	case dto.BulkGetAccountsOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]interface{}, dto.BulkGetAccountsOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodPost,
		"bulk/items",
		params,
		nil,
	)
}

func (m *Market) EditAccount(itemID int, query any, body any) (*dto.Response, error) {
	switch v := query.(type) {
	case nil:
	case map[string]string:
	case dto.EditAccountOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип query: %T | Поддерживаемые: map[string]string, dto.EditAccountOptional",
			v,
		)
	}

	switch v := body.(type) {
	case nil:
	case map[string]interface{}:
	case dto.EditAccountBodyOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип body: %T | Поддерживаемые: map[string]interface{}, dto.EditAccountBodyOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodPut,
		fmt.Sprintf("%d/edit", itemID),
		query,
		body,
	)
}

func (m *Market) DeleteAccount(itemId int, reason string) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodDelete,
		fmt.Sprintf("%d", itemId),
		map[string]string{"reason": reason},
		nil,
	)
}

func (m *Market) BumpAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/bump", itemId),
		nil,
		nil,
	)
}

func (m *Market) OpenAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/open", itemId),
		nil,
		nil,
	)
}

func (m *Market) CloseAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/close", itemId),
		nil,
		nil,
	)
}

func (m *Market) CreateArbitrage(itemId int, postBody string) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/claims", itemId),
		nil,
		map[string]interface{}{"post_body": postBody},
	)
}

func (m *Market) EditNote(itemId int, text string) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/note-save", itemId),
		nil,
		map[string]interface{}{"text": text},
	)
}

func (m *Market) AccountImage(itemID int, imageType string) (*dto.Response, error) {
	switch imageType {
	case dto.AccountImageTypeSkins,
		dto.AccountImageTypePickaxes,
		dto.AccountImageTypeDances,
		dto.AccountImageTypeGliders,
		dto.AccountImageTypeWeapons,
		dto.AccountImageTypeAgents,
		dto.AccountImageTypeBuddies:
	default:
		return nil, fmt.Errorf("неподдерживаемый тип: %q", imageType)
	}

	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/image", itemID),
		map[string]string{
			"type": imageType,
		},
		nil,
	)
}

func (m *Market) EmailLetters(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case nil:
	case map[string]string:
	case dto.EmailLettersOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.GetEmailLettersOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodGet,
		"letters2",
		params,
		nil,
	)
}

func (m *Market) EmailCode(params any) (*dto.Response, error) {
	switch v := params.(type) {
	case nil:
	case map[string]string:
	case dto.EmailCodeOptional:
	default:
		return nil, fmt.Errorf(
			"неподдерживаемый тип params: %T | Поддерживаемые: map[string]string, dto.GetEmailCodeOptional",
			v,
		)
	}

	return m.doRequest(
		resty.MethodGet,
		"email-code",
		params,
		nil,
	)
}

func (m *Market) CheckGuarantee(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/check-guarantee", itemId),
		nil,
		nil,
	)
}

func (m *Market) CancelGuarantee(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/refuse-guarantee", itemId),
		nil,
		nil,
	)
}

func (m *Market) ChangePassword(itemId int, cancel int) (*dto.Response, error) {
	if cancel != 1 {
		return nil, fmt.Errorf("cancel может быть значением 1")
	}

	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/refuse-guarantee", itemId),
		nil,
		nil,
	)
}

func (m *Market) TempEmailPassword(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/temp-email-password", itemId),
		nil,
		nil,
	)
}

func (m *Market) TagAccount(itemId, tagId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/tag", itemId),
		map[string]int{"tag_id": tagId},
		nil,
	)
}

func (m *Market) UntagAccount(itemId, tagId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodDelete,
		fmt.Sprintf("%d/tag", itemId),
		map[string]int{"tag_id": tagId},
		nil,
	)
}

func (m *Market) Favorite(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/star", itemId),
		nil,
		nil,
	)
}

func (m *Market) UnFavorite(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodDelete,
		fmt.Sprintf("%d/star", itemId),
		nil,
		nil,
	)
}

func (m *Market) StickAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/stick", itemId),
		nil,
		nil,
	)
}

func (m *Market) UnStickAccount(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodDelete,
		fmt.Sprintf("%d/stick", itemId),
		nil,
		nil,
	)
}

func (m *Market) ChangeAccountOwner(itemId int, username, secretAnswer string) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodPost,
		fmt.Sprintf("%d/change-owner", itemId),
		map[string]string{"username": username, "secret_answer": secretAnswer},
		nil,
	)
}

func (m *Market) AIPrice(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/ai-price", itemId),
		nil,
		nil,
	)
}

func (m *Market) AutoBuyPrice(itemId int) (*dto.Response, error) {
	return m.doRequest(
		resty.MethodGet,
		fmt.Sprintf("%d/auto-buy-price", itemId),
		nil,
		nil,
	)
}
