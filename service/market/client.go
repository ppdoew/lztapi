package market

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/ppdoew/lztapi/dto"
	"github.com/ppdoew/lztapi/utils"
	"net/url"
)

func (m *Market) doRequest(
	method string,
	path string,
	value any, // query-параметры: map[string]string, url.Values или struct с тегами `url:""`
	body any, // тело: map[string]interface{} или struct с тегами `json:""`
) (*dto.Response, error) {
	req := m.cli.Client.R()

	if value != nil {
		switch v := value.(type) {
		case map[string]string:
			req = req.SetQueryParams(v)

		case url.Values:
			req = req.SetQueryParamsFromValues(v)

		default:
			vals, err := query.Values(v)
			if err != nil {
				return nil, fmt.Errorf("не удалось сериализовать query-параметры: %w", err)
			}
			req = req.SetQueryParamsFromValues(vals)
		}
	}

	if body != nil {
		req = req.SetHeader("Content-Type", "application/json")

		switch v := body.(type) {
		case map[string]interface{}:
			req = req.SetBody(v)

		default:
			req = req.SetBody(v)
		}
	}

	var (
		resp *resty.Response
		err  error
	)
	switch method {
	case resty.MethodGet:
		resp, err = req.Get(path)
	case resty.MethodPost:
		resp, err = req.Post(path)
	case resty.MethodPut:
		resp, err = req.Put(path)
	case resty.MethodDelete:
		resp, err = req.Delete(path)
	default:
		resp, err = req.Execute(method, path)
	}
	if err != nil {
		return nil, err
	}

	return &dto.Response{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Text:       resp.String(),
		Json:       utils.GetJson(resp),
	}, nil
}
