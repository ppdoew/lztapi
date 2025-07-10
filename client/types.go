package client

import (
	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

type LztClient struct {
	Client  *resty.Client
	Limiter *rate.Limiter
	baseURL string
}
