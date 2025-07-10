package market

import (
	"github.com/go-resty/resty/v2"
	"github.com/ppdoew/lztapi/dto"
)

func (m *Market) Batch(request []dto.BatchJob) (*dto.Response, error) {
	return m.doRequest(resty.MethodPost, "batch", nil, request)
}
