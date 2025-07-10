package utils

import (
	"LztApi/lztapi/errors"
	"encoding/json"
	"github.com/go-resty/resty/v2"
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
		return map[string]any{"error": []string{"undefined"}}
	}

	return result
}
