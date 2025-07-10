package client

import (
	"LztApi/lztapi/config"
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

func NewSession(token string) *LztClient {
	limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 1)

	client := resty.New().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		}).
		SetTimeout(config.DefaultTimeout).
		SetRetryCount(config.DefaultRetryCount).
		SetRetryWaitTime(config.DefaultRetryWait).
		SetRetryMaxWaitTime(config.DefaultRetryMaxWait).
		SetCloseConnection(true)

	lzt := &LztClient{
		Client:  client,
		Limiter: limiter,
		baseURL: config.DefaultBaseURL,
	}

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		if !strings.HasPrefix(r.URL, "http://") && !strings.HasPrefix(r.URL, "https://") {
			r.URL = lzt.baseURL + r.URL
		}
		return lzt.Limiter.Wait(context.Background())
	})

	return lzt
}

func (l *LztClient) SetProxy(proxyAddr string) error {
	if proxyAddr == "" {
		l.Client.SetProxy("")
		return nil
	}
	u, err := url.Parse(proxyAddr)
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "socks5" {
		return errors.New("недопустимая схема прокси: только http, https, ")
	}
	l.Client.SetProxy(proxyAddr)
	return nil
}

func (l *LztClient) SetLimiter(newLimiter *rate.Limiter) {
	if newLimiter != nil {
		l.Limiter = newLimiter
	}
}

func (l *LztClient) SetTimeout(d time.Duration) {
	l.Client.SetTimeout(d)
}

func (l *LztClient) SetRetryCount(count int) {
	l.Client.SetRetryCount(count)
}

func (l *LztClient) SetRetryWaitTime(d time.Duration) {
	l.Client.SetRetryWaitTime(d)
}

func (l *LztClient) SetRetryMaxWaitTime(d time.Duration) {
	l.Client.SetRetryMaxWaitTime(d)
}
