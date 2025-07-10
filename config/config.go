package config

import "time"

const (
	DefaultBaseURL      = "https://prod-api.lzt.market/"
	DefaultTimeout      = 120 * time.Second
	DefaultRetryCount   = 3
	DefaultRetryWait    = 3 * time.Second
	DefaultRetryMaxWait = 3 * time.Second
)
