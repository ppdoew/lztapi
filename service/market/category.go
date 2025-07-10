package market

import (
	"LztApi/lztapi/dto"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func (m *Market) CategorySearch(category dto.Category, params map[string]string) (*dto.Response, error) {
	if !dto.IsValidCategory(category) {
		return nil, fmt.Errorf("недоступна категория %q", category)
	}
	return m.doRequest(resty.MethodGet, string(category), params, nil)
}

func (m *Market) CategoryGames(category dto.Category) (*dto.Response, error) {
	if category == "" || !dto.IsValidCategory(category) {
		return nil, fmt.Errorf("недоступна категория %q", category)
	}
	path := fmt.Sprintf("%s/games", string(category))
	return m.doRequest(resty.MethodGet, path, nil, nil)
}

func (m *Market) CategorySearchParams(category dto.Category) (*dto.Response, error) {
	if category == "" || !dto.IsValidCategory(category) {
		return nil, fmt.Errorf("недоступна категория %q", category)
	}
	path := fmt.Sprintf("%s/params", string(category))
	return m.doRequest(resty.MethodGet, path, nil, nil)
}

func (m *Market) Categories(topQueries ...bool) (*dto.Response, error) {
	var params map[string]string
	if len(topQueries) > 0 {
		params = map[string]string{
			"top_queries": strconv.FormatBool(topQueries[0]),
		}
	}
	return m.doRequest(resty.MethodGet, "category", params, nil)
}
