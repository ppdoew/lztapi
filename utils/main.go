package utils

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/errors"
)

func GetJson(response *resty.Response) map[string]any {
	var result map[string]any

	status := response.StatusCode()

	msg, ok := errors.ErrorMessages[status]
	if ok {
		return map[string]any{
			"error": msg,
		}
	}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return map[string]any{"errors": []string{"undefined"}}
	}

	return result
}
