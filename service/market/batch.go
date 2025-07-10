package market

import (
	"LztApi/lztapi/dto"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (m *Market) Batch(request []dto.BatchJob) (*dto.Response, error) {

	fmt.Println(request)
	return m.doRequest(resty.MethodPost, "batch", nil, request)
}
