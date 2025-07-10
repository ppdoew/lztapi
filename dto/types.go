package dto

type Response struct {
	StatusCode int                    `json:"status_code"`
	Body       []byte                 `json:"body"`
	Text       string                 `json:"text"`
	Json       map[string]interface{} `json:"json"`
}
